package engine

type EffectType string

type Effecter interface {
	GetType() EffectType
	Apply()
	Cancel()
}

type EffectListener interface {
	onStack(Effecter)
	onPop(Effecter)
	onCancel(Effecter)
	onApply(Effecter)
}

type EffectStack struct {
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

func (es *EffectStack) AddListener(l EffectListener) {
	for EffectType := range es.listeners {
		es.listeners[EffectType] = append(es.listeners[EffectType], l)
	}
}

func (es *EffectStack) Resolve() {
	for len(es.Effects) != 0 {
		e := es.PopEffect()
		e.Apply()
		if es.listeners != nil {
			for _, l := range es.listeners[e.GetType()] {
				l.onApply(e)
			}
		}
	}
}
