package engine

// Operator is an interface that defines methods for evaluating conditions
// and retrieving child conditions. Implementers of this interface should
// provide logic for the Evaluate method to determine if a condition is met,
// and the GetChildren method to return any sub-conditions.

type Operator interface {
	Evaluate(OperatorNode) bool
}

// A single rule
type Ruler interface {
	Apply(Rule) bool
}

type OperatorNode struct {
	Operator string `json:"operator"`
	Children []OperatorNode
}

type ActionProcessor interface {
	Preprocess(game Game, entity EntityID, values []KeyValue)
	Process(game Game, entity EntityID, values []KeyValue)
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Rule struct {
	Name      string     `json:"name"`
	Condition any        `json:"condition"`
	Action    []KeyValue `json:"action"`
}

type RuleSection struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type RuleList struct {
	Sections []RuleSection `json:"sections"`
}

type OperatorType string

type OperatorRegistry struct {
	Operators map[OperatorType]Operator
}

type SectionType string

type RuleRegistry struct {
	sections  map[SectionType]Ruler
	operators OperatorRegistry
}

func (r *RuleRegistry) RegisterRuleSection(section SectionType, rule Rule) {
	if r.sections == nil {
		r.sections = make(map[SectionType]Ruler)
	}
	if ruler, exists := r.sections[section]; exists {
		ruler.Apply(rule)
	}
}

func (r *RuleRegistry) GetRuleSection(section SectionType) Ruler {
	if ruler, exists := r.sections[section]; exists {
		return ruler
	}
	return nil
}

func (r *RuleRegistry) GetOperator(operator OperatorType) Operator {
	if op, exists := r.operators.Operators[operator]; exists {
		return op
	}
	return nil
}

func (r *RuleRegistry) RegisterOperator(operator OperatorType, op Operator) {
	if r.operators.Operators == nil {
		r.operators.Operators = make(map[OperatorType]Operator)
	}
	r.operators.Operators[operator] = op
}
