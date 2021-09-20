package mapper

import (
	"fmt"
	"strings"
	"sync"

	"github.com/fantadeltaalpha/flogo/extension/activity/maptree/model"
	"github.com/fantadeltaalpha/flogo/extension/activity/maptree/model/ope"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Mapper struct{

}


func New() *Mapper {
	return &Mapper{}
}

func (m *Mapper) ToProduct(data []byte) (*model.Products,error){
	resp:= ope.GetOfferResponse{}
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return nil,err
	}

	poList := model.Products{Products: make([]*model.Product, 0)}
	mrMap := make(map[string]*model.Product, 0)

	wg := sync.WaitGroup{}
	mPO  := sync.Mutex{}
	mMR  := sync.Mutex{}

	for _,e := range resp.EligibleProducts.EligibleProduct{
		wg.Add(1)
		go func(e ope.EligibleProduct){
			defer wg.Done()
			prod := model.Product{
				ID: e.ProductInformation.Product.ProductID,
				Name: e.ProductInformation.Product.ProductName,
				Type: e.ProductInformation.Product.ProductType,
			}
			prod.PCO = make([]string, 0)
			prod.Attribute = make(map[string]string)
			for _,char := range e.ProductInformation.Characteristic {
				if char.Description  ==  "ProductComprisedOf"{
					prod.PCO = append(prod.PCO, char.Name)
				}else if char.Description  ==  "Characteristic" {
					if len(char.Values.Value) >0 {
						prod.Attribute[strings.ToLower(char.Name)] = char.Values.Value[0].Value
					}
				}
			}
			switch e.ProductInformation.Product.ProductType {
			case "PO":
				prod.ProductComprisedOf = make([]model.Product, 0)
				mPO.Lock()
				poList.Products= append(poList.Products, &prod)
				mPO.Unlock()
			case "MR":
				mMR.Lock()
				mrMap[e.ProductInformation.Product.ProductID] = &prod
				mMR.Unlock()	
			}
			
		}(e)
	}
	wg.Wait()

	for _, po := range poList.Products {
		for _, pco := range po.PCO {
			if v,ok := mrMap[pco];ok{
				po.ProductComprisedOf = append(po.ProductComprisedOf, *v)
			}else{
				fmt.Printf("Not Found %v\n", pco)
			}
		}
	}

	return &poList,nil
}	