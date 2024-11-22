package engine

type EffectType string

type Effecter interface {
	GetType() EffectType
	// Apply(*EffectStack)
	// Cancel()
}

type EffectListener interface {
	OnStack(Effecter, *EffectStack)
	OnPop(Effecter, *EffectStack)
	OnCancel(Effecter, *EffectStack)
	OnApply(Effecter, *EffectStack)
}

type EffectStack struct {
	Game      *Game
	Effects   []Effecter
	listeners map[EffectType][]EffectListener
}

func (es *EffectStack) StackEffect(e Effecter) {
	es.Effects = append(es.Effects, e)
}

func (es *EffectStack) PopEffect() Effecter {
	if len(es.Effects) == 0 {
		return nil
	}
	e := es.Effects[len(es.Effects)-1]
	es.Effects = es.Effects[0 : len(es.Effects)-1]
	return e
}

func (es *EffectStack) CancelEffect(e Effecter) {
	if len(es.Effects) == 0 {
		return
	}
	for i, l := range es.Effects {
		if l == e {
			if es.listeners != nil {
				for _, l := range es.listeners[e.GetType()] {
					l.OnCancel(e, es)
				}
			}
			es.Effects = append(es.Effects[:i], es.Effects[i+1:]...)
			return
		}
	}
}

func (es *EffectStack) AddListener(et EffectType, l EffectListener) {
	if es.listeners == nil {
		es.listeners = make(map[EffectType][]EffectListener, 100)
	}
	es.listeners[et] = append(es.listeners[et], l)
}

func (es *EffectStack) Resolve() {
	for len(es.Effects) != 0 {
		e := es.PopEffect()
		//e.Apply(es)
		if es.listeners != nil {
			for _, l := range es.listeners[e.GetType()] {
				l.OnApply(e, es)
			}
		}
	}
	//fmt.Println("Number of remaining effects:", len(es.Effects))
}
