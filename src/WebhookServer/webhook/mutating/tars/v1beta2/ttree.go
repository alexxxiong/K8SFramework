package v1beta2

import (
	"fmt"
	k8sAdmissionV1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/util/json"
	tarsV1beta2 "k8s.tars.io/apis/tars/v1beta2"
	tarsTool "k8s.tars.io/tool"
	"tarswebhook/webhook/lister"
	"tarswebhook/webhook/mutating"
)

func mutatingCreateTTree(listers *lister.Listers, requestAdmissionView *k8sAdmissionV1.AdmissionReview) ([]byte, error) {
	newTTree := &tarsV1beta2.TTree{}
	_ = json.Unmarshal(requestAdmissionView.Request.Object.Raw, newTTree)

	businessMap := make(map[string]interface{}, len(newTTree.Businesses))
	for _, business := range newTTree.Businesses {
		businessMap[business.Name] = nil
	}

	var jsonPatch tarsTool.JsonPatch

	for i, app := range newTTree.Apps {
		if app.BusinessRef != "" {
			if _, ok := businessMap[app.BusinessRef]; !ok {
				jsonPatch = append(jsonPatch, tarsTool.JsonPatchItem{
					OP:    tarsTool.JsonPatchReplace,
					Path:  fmt.Sprintf("/apps/%d/businessRef", i),
					Value: "",
				})
			}
		}
	}

	if jsonPatch != nil {
		return json.Marshal(jsonPatch)
	}
	return nil, nil
}

func mutatingUpdateTTree(listers *lister.Listers, requestAdmissionView *k8sAdmissionV1.AdmissionReview) ([]byte, error) {
	return mutatingCreateTTree(listers, requestAdmissionView)
}

func init() {
	gvr := tarsV1beta2.SchemeGroupVersion.WithResource("ttrees")
	mutating.Registry(k8sAdmissionV1.Create, &gvr, mutatingCreateTTree)
	mutating.Registry(k8sAdmissionV1.Update, &gvr, mutatingUpdateTTree)
}
