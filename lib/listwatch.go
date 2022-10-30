package lib

import (
	"github.com/kevinlisr/handlers"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"log"
)

func Watch(fact informers.SharedInformerFactory,g, v, r string)  cache.SharedIndexInformer {
	if g == "core"{
		g = ""
	}
	cmGVR := schema.GroupVersionResource{
		Group:    g,
		Version:  v,
		Resource: r,
	}
	cmInformer, err := fact.ForResource(cmGVR)
	if err != nil {
		log.Fatalln(err)
	}


	return cmInformer.Informer()
}



func CmWatch(fact informers.SharedInformerFactory,g,v, r string)  {
	cmGVR := schema.GroupVersionResource{
		Group:    g,
		Version:  v,
		Resource: r,
	}
	cmInformer, err := fact.ForResource(cmGVR)
	if err != nil {
		log.Fatalln(err)
	}
	cmInformer.Informer().AddEventHandler(&handlers.CmHandler{})
}

func DpWatch(fact informers.SharedInformerFactory,gv, rs string)  {
	cmGVR := schema.GroupVersionResource{
		Group:    "apps",
		Version:  gv,
		Resource: rs,
	}
	cmInformer, err := fact.ForResource(cmGVR)
	if err != nil {
		log.Fatalln(err)
	}
	cmInformer.Informer().AddEventHandler(&handlers.CmHandler{})
}

func Start(factory informers.SharedInformerFactory)  {
	factory.Start(wait.NeverStop)
}