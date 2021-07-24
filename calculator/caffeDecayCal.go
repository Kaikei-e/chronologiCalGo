package calculator

import (
	drinkvalidator "caffecalgo/drinkValidator"
	"log"
	"time"
)

type caffeineDecay struct{
	DecayCaffe float64
	DecayTime time.Time
}

type CaffeineDecays struct{
	CaffeLists []caffeineDecay
}


func CaffeDecayCals(caffeStructs []drinkvalidator.CaffeLogger) CaffeineDecays {
	var caffeDcays CaffeineDecays
	var isTmaxed bool

	listLength := len(caffeStructs)
	timeLimit := []time.Duration{}
	j := 0


	for i := 0; i < listLength -1; i++ {
		j += 1
		periodOfTime := caffeStructs[j].Datetime.Sub(caffeStructs[i].Datetime)

		timeLimit = append(timeLimit, periodOfTime)
	}

	log.Println("listLength")
	log.Println(listLength)
	log.Println("timeLimit")
	log.Println(timeLimit)

	if listLength == 1{
		calMethodSimple(caffeStructs[0])

	}

	for i := 0; i < listLength - 1; i++ {

		j += 1
		if j > listLength {
			break
		}

		if (listLength == 2){

		}else{
			caffeDcays, isTmaxed = calTmax(caffeStructs[i], timeLimit[i + 1])

			if isTmaxed {
				timeLast := caffeDcays.CaffeLists[len(caffeDcays.CaffeLists) - 1].DecayTime
				calDecay(caffeStructs[i], caffeDcays, timeLast)
				log.Println("isTmax check")


				break
			}
		}

	}

	return caffeDcays
}

func calTmax(caffeStruct drinkvalidator.CaffeLogger, periodOfTime time.Duration) (CaffeineDecays, bool){
	var caffeDecays CaffeineDecays
	minutes := int64(periodOfTime / time.Minute)
	const TmaxVar = 1.1333
	const startTmax = 1
	var caffeDe caffeineDecay
	caffeDe.DecayTime = caffeStruct.Datetime
	var amountOfCaffeine int

	var isTmax bool

	if caffeStruct.Method == 2 {
		amountOfCaffeine = caffeStruct.Amount / 100 * caffeStruct.CaffeineMg
	}else{
		amountOfCaffeine = caffeStruct.CaffeineMg
	}

	for i := 0; i < int(minutes); i++ {
		if caffeDe.DecayCaffe > float64(amountOfCaffeine) {
			isTmax = true
			break
		}
		if caffeStruct.Datetime.After(caffeDe.DecayTime) {
			isTmax = false
			break
		}

		caffeDe.DecayCaffe += startTmax * TmaxVar
		caffeDe.DecayTime = caffeDe.DecayTime.Add(time.Duration(1) * time.Minute)

		caffeDecays.CaffeLists = append(caffeDecays.CaffeLists, caffeDe)
		log.Println(i)
		log.Println(caffeDe.DecayCaffe)
		log.Println(caffeDe.DecayTime)
		log.Println(caffeStruct.Datetime)

	}

	return caffeDecays, isTmax
}

func calDecay( caffeSt drinkvalidator.CaffeLogger, caffeDcays CaffeineDecays, timeLast time.Time){
	const decayRate = 0.99807

	periodOfTime := timeLast.Sub(caffeSt.Datetime)
	minutes := int64(periodOfTime / time.Minute)
	log.Println(minutes)
	caffeDe := caffeDcays.CaffeLists[len(caffeDcays.CaffeLists) - 1]

	amountOfCaffeine := caffeDe.DecayCaffe

	for i := 0; i < int(minutes); i ++ {
		if amountOfCaffeine < 5 {
			break
		}

		if caffeSt.Datetime.After(caffeDe.DecayTime) {
			break
		}
		amountOfCaffeine *= decayRate
		log.Println(amountOfCaffeine)
	}





}

func calMethodSimple(caffeStruct drinkvalidator.CaffeLogger) CaffeineDecays {
	var caffeDecays CaffeineDecays
	const TmaxVar = 1.3397
	const startTmax = 1
	var caffeDe caffeineDecay
	caffeDe.DecayTime = caffeStruct.Datetime
	var amountOfCaffeine float64
	const maxCount = 10000


	if caffeStruct.Method == 2 {
		amountOfCaffeine = float64(caffeStruct.Amount) / 100 * float64(caffeStruct.CaffeineMg)
	}else{
		amountOfCaffeine = float64(caffeStruct.CaffeineMg)
	}

	for i := 0; i < maxCount; i++ {
		if caffeDe.DecayCaffe > float64(amountOfCaffeine) {
			break
		}

		caffeDe.DecayCaffe += startTmax * TmaxVar
		caffeDe.DecayTime = caffeDe.DecayTime.Add(time.Duration(5) * time.Minute)

		caffeDecays.CaffeLists = append(caffeDecays.CaffeLists, caffeDe)


	}

	const decayRate = 0.98832
	const footstep = 5

	for i := 0; i < maxCount; i ++ {
		if amountOfCaffeine < footstep {
			break
		}

		amountOfCaffeine *= decayRate

		caffeDe.DecayCaffe = amountOfCaffeine
		caffeDe.DecayTime = caffeDe.DecayTime.Add(time.Duration(5) * time.Minute)

		caffeDecays.CaffeLists = append(caffeDecays.CaffeLists, caffeDe)


	}

	log.Println(len(caffeDecays.CaffeLists))


	return caffeDecays
}
