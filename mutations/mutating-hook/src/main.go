package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// mutateDeployment updates the replicas of a Deployment from 1 to 2
func mutateDeployment(review admissionv1.AdmissionReview) *admissionv1.AdmissionResponse {
	var deployment appsv1.Deployment
	if err := json.Unmarshal(review.Request.Object.Raw, &deployment); err != nil {
		log.Printf("Error unmarshalling Deployment: %v", err)
		return toAdmissionResponse(err)
	}

	// Only modify if replicas == 1
	if deployment.Spec.Replicas != nil && *deployment.Spec.Replicas == 1 {
		var patchOps []map[string]interface{}
		patchOps = append(patchOps, map[string]interface{}{
			"op":    "replace",
			"path":  "/spec/replicas",
			"value": 2,
		})

		patchBytes, err := json.Marshal(patchOps)
		if err != nil {
			return toAdmissionResponse(err)
		}

		log.Printf("Mutating deployment %s/%s replicas from 1 → 2", deployment.Namespace, deployment.Name)
		return &admissionv1.AdmissionResponse{
			Allowed:   true,
			Patch:     patchBytes,
			PatchType: func() *admissionv1.PatchType { pt := admissionv1.PatchTypeJSONPatch; return &pt }(),
		}
	}

	// If no change, just allow
	return &admissionv1.AdmissionResponse{Allowed: true}
}

func admitHandler(w http.ResponseWriter, r *http.Request) {
	var review admissionv1.AdmissionReview
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := admissionv1.AdmissionReview{
		TypeMeta: review.TypeMeta,
		Response: mutateDeployment(review),
	}
	response.Response.UID = review.Request.UID

	respBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}

func toAdmissionResponse(err error) *admissionv1.AdmissionResponse {
	return &admissionv1.AdmissionResponse{
		Result: &metav1.Status{
			Message: err.Error(),
		},
	}
}

func main() {
	var certFile, keyFile string
	flag.StringVar(&certFile, "tls-cert-file", "/tls/tls.crt", "TLS certificate file")
	flag.StringVar(&keyFile, "tls-private-key-file", "/tls/tls.key", "TLS private key file")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/mutate", admitHandler)

	server := &http.Server{
		Addr:      ":8443",
		Handler:   mux,
		TLSConfig: &tls.Config{},
	}

	log.Println("Starting mutating webhook server on :8443 ...")
	if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
