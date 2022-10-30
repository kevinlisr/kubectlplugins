package handlers

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
)


type CmHandler struct {

}

func (this *CmHandler) OnAdd(obj interface{})  {
	fmt.Println("add",obj.(*v1.ConfigMap).Name)
}

func (this *CmHandler) OnUpdate(oldObj,newObj interface{})  {
	//fmt.Println("update",newObj.(*v1.ConfigMap).Name,newObj.(*v1.ConfigMap))
}

func (this *CmHandler) OnDelete(obj interface{})  {
	fmt.Println("delete",obj.(*v1.ConfigMap).Name)
}

//type CmHandlerNew struct {
//
//}
//
//func (this *CmHandlerNew) OnAdd(obj interface{})  {
//	//fmt.Println("add",obj.(*v1.ConfigMap).Name)
//}
//
//func (this *CmHandlerNew) OnUpdate(oldObj,newObj interface{})  {
//	//fmt.Println("update",newObj.(*v1.ConfigMap).Name,newObj.(*v1.ConfigMap))
//}
//
//func (this *CmHandlerNew) OnDelete(obj interface{})  {
//	//fmt.Println("delete",obj.(*v1.ConfigMap).Name)
//}