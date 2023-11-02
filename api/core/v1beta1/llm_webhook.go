/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"github.com/DataTunerX/meta-server/logging"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var llmlog = logging.Logger.WithName("llm-resource")

func (r *LLM) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-core-datatunerx-io-v1beta1-llm,mutating=true,failurePolicy=fail,sideEffects=None,groups=core.datatunerx.io,resources=llms,verbs=create;update,versions=v1beta1,name=mllm.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &LLM{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *LLM) Default() {
	llmlog.Info("default", "name", r.Name)

}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-core-datatunerx-io-v1beta1-llm,mutating=false,failurePolicy=fail,sideEffects=None,groups=core.datatunerx.io,resources=llms,verbs=create;update,versions=v1beta1,name=vllm.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &LLM{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *LLM) ValidateCreate() error {
	llmlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *LLM) ValidateUpdate(old runtime.Object) error {
	llmlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *LLM) ValidateDelete() error {
	llmlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}