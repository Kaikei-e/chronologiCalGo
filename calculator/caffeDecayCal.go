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

type caffeineDecays struct{
	caffeLists []caffeineDecay
}


func CaffeDecayCals(caffeStructs []drinkvalidator.CaffeLogger){
	var caffeDcays caffeineDecays
	var isTmaxed bool

	listLength := len(caffeStructs)
	timeLimit := []time.Duration{}
	j := 0


	for i := 0; i < listLength -1; i++ {
		j += 1
		periodOfTime := caffeStructs[j].Datetime.Sub(caffeStructs[i].Datetime)

		timeLimit = append(timeLimit, periodOfTime)
	}

	log.Println(timeLimit)


	for i := 0; i < listLength - 1; i++ {
		j += 1
		if j > listLength {
			break
		}

		if i == listLength - 2{
			calMethodSimple(caffeStructs[i])
		}else{

			if caffeStructs[i].Method == 1 {
				caffeDcays, isTmaxed = calTmax(caffeStructs[i], timeLimit[i + 1])

				if isTmaxed {
					timeLast := caffeDcays.caffeLists[len(caffeDcays.caffeLists) - 1].decayTime
					calDecay(caffeStructs[i], caffeDcays, timeLast)
					log.Println("check")


					break
				}else{
				}
			}else if (caffeStructs[i].Method == 2){
				//calDecay(caffeLogs[i], timeDuration[i + 1])
			}
		}

	}
}

func calTmax(caffeStruct drinkvalidator.CaffeLogger, periodOfTime time.Duration) (caffeineDecays, bool){
	var caffeDecays caffeineDecays
	minutes := int64(periodOfTime / time.Minute)
	const TmaxVar = 1.1333
	const startTmax = 1
	var caffeDe caffeineDecay
	caffeDe.decayTime = caffeStruct.Datetime
	var amountOfCaffeine int

	var isTmax bool

	if caffeStruct.Method == 2 {
		amountOfCaffeine = caffeStruct.Amount / 100 * caffeStruct.CaffeineMg
	}else{
		amountOfCaffeine = caffeStruct.CaffeineMg
	}

	for i := 0; i < int(minutes); i++ {
		if caffeDe.decayCaffe > float64(amountOfCaffeine) {
			isTmax = true
			break
		}
		if caffeStruct.Datetime.After(caffeDe.decayTime) {
			isTmax = false
			break
		}

		caffeDe.decayCaffe += startTmax * TmaxVar
		caffeDe.decayTime = caffeDe.decayTime.Add(time.Duration(1) * time.Minute)

		caffeDecays.caffeLists = append(caffeDecays.caffeLists, caffeDe)
		log.Println(i)
		log.Println(caffeDe.decayCaffe)
		log.Println(caffeDe.decayTime)
		log.Println(caffeStruct.Datetime)

	}

	return caffeDecays, isTmax
}

func calDecay( caffeSt drinkvalidator.CaffeLogger, caffeDcays caffeineDecays, timeLast time.Time){
	const decayRate = 0.99807

	periodOfTime := timeLast.Sub(caffeSt.Datetime)
	minutes := int64(periodOfTime / time.Minute)
	log.Println(minutes)
	caffeDe := caffeDcays.caffeLists[len(caffeDcays.caffeLists) - 1]

	amountOfCaffeine := caffeDe.decayCaffe

	for i := 0; i < int(minutes); i ++ {
		if amountOfCaffeine < 5 {
			break
		}

		if caffeSt.Datetime.After(caffeDe.decayTime) {
			break
		}
		amountOfCaffeine *= decayRate
		log.Println(amountOfCaffeine)
	}



}

func calMethodSimple(caffeStruct drinkvalidator.CaffeLogger){

}
