package common

type DeployType string

const (
	Develop    DeployType = "dev"
	Production DeployType = "prod"
	Test       DeployType = "test"
)

func (t DeployType) IsDevelop() bool {
	return t == Develop
}

func (t DeployType) IsProduction() bool {
	return t == Production
}

func (t DeployType) IsTest() bool {
	return t == Test
}
