package calculator

import (
	drinkvalidator "caffecalgo/drinkValidator"
	"log"
	"time"
)

type caffeineDecay struct{
	decayCaffe float64
	decayTime time.Time 
}


func CaffeDecayCals(caffeLogs []drinkvalidator.CaffeLogger){
	//caffeineDecays := []caffeineDecay{}
	listLength := len(caffeLogs)
	timeDuration := []time.Duration{}
	j := 0


	for i := 0; i < listLength -1; i++ {
		j += 1
		periodOfTime := caffeLogs[j].Datetime.Sub(caffeLogs[i].Datetime)

		timeDuration = append(timeDuration, periodOfTime)
	}

	log.Println(timeDuration)


	for i := 0; i < listLength - 1; i++ {
		j += 1
		if j > listLength {
			break
		}

		if i == listLength - 2{
			calMethodSimple(caffeLogs[i])
		}else{
			
			if caffeLogs[i].Method == 1 {
				calMethod1(caffeLogs[i], timeDuration[i + 1])
			}else if (caffeLogs[i].Method == 2){
				calMethod2(caffeLogs[i], timeDuration[i + 1])
			}			
		}

	}
}

func calMethod1(caffeStruct drinkvalidator.CaffeLogger, periodOfTime time.Duration){
	caffeineDecays := []caffeineDecay{}
	minutes := int64(periodOfTime / time.Minute)
	TmaxVar := 1.1333
	var caffeDe caffeineDecay
	caffeDe.decayTime = caffeStruct.Datetime

	for i := 0; i < int(minutes); i++ {

		caffeDe.decayCaffe += 1 * TmaxVar
		caffeDe.decayTime.Add(time.Duration(1) * time.Minute)

		caffeineDecays = append(caffeineDecays, caffeDe)
		//log.Println(caffeDe)
	}
}

func calMethod2(caffeStruct drinkvalidator.CaffeLogger, periodOfTime time.Duration){



}

func calMethodSimple(caffeStruct drinkvalidator.CaffeLogger){

}