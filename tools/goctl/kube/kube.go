package kube

import (
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/urfave/cli"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	category           = "kube"
	deployTemplateFile = "deployment.tpl"
	jobTemplateFile    = "job.tpl"
	basePort           = 30000
	portLimit          = 32767
)

// Deployment describes the k8s deployment yaml
type Deployment struct {
	Name              string
	Namespace         string
	Image             string
	Secret            string
	Domain            string
	EnableGRPCIngress bool   //开始环境启用rpc，其他环境不启用对外的RPC服务
	IngressPrefix     string //用于HTTP ingress的前缀
	Replicas          int
	Revisions         int
	Port              int
	NodePort          int
	UseNodePort       bool
	RequestCpu        int
	RequestMem        int
	LimitCpu          int
	LimitMem          int
	MinReplicas       int
	MaxReplicas       int
	ServiceAccount    string
	EnableTls         bool
	Env               string
	EnableWebsocket   bool
	IsNotUserCenter   bool
}

// DeploymentCommand is used to generate the kubernetes deployment yaml files.
func DeploymentCommand(c *cli.Context) error {
	nodePort := c.Int("nodePort")
	home := c.String("home")
	remote := c.String("remote")
	if len(remote) > 0 {
		repo, _ := util.CloneIntoGitHome(remote)
		if len(repo) > 0 {
			home = repo
		}
	}

	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}
	domain := c.String("domain")
	if domain == "" {
		domain = "api.flyele.vip"
	}
	// 0 to disable the nodePort type
	if nodePort != 0 && (nodePort < basePort || nodePort > portLimit) {
		return errors.New("nodePort should be between 30000 and 32767")
	}

	var (
		env          = c.String("env")
		kubeTemplate = deploymentTemplate
	)
	if env == "k3sprod" {
		kubeTemplate = k3sDeploymentTemplate
	}
	if c.String("name") == "push-gateway" {
		kubeTemplate = statefullsetTemplate
	}

	text, err := pathx.LoadTemplate(category, deployTemplateFile, kubeTemplate)
	if err != nil {
		return err
	}

	out, err := pathx.CreateIfNotExist(c.String("o"))
	if err != nil {
		return err
	}
	defer out.Close()
	funcMap := template.FuncMap{
		"getProfileDomain": func(domain string) string {
			return strings.Replace(domain, "api", "profile", -1)
		},
		"getAppKind": func(serviceName string) string {
			if serviceName == "push-gateway" {
				return "StatefulSet"
			}
			return "Deployment"
		},
	}
	t := template.Must(template.New("deploymentTemplate").Funcs(funcMap).Parse(text))

	err = t.Execute(out, Deployment{
		Domain:            domain,
		EnableWebsocket:   c.String("name") == "push-gateway",
		EnableGRPCIngress: c.Bool("enable-grpc-ingres"),
		IngressPrefix:     c.String("ingress-prefix"),
		EnableTls:         c.Bool("enable-tls"),
		Env:               c.String("env"),
		Name:              c.String("name"),
		Namespace:         c.String("namespace"),
		Image:             c.String("image"),
		Secret:            c.String("secret"),
		Replicas:          c.Int("replicas"),
		Revisions:         c.Int("revisions"),
		Port:              c.Int("port"),
		NodePort:          nodePort,
		UseNodePort:       nodePort > 0,
		RequestCpu:        c.Int("requestCpu"),
		RequestMem:        c.Int("requestMem"),
		LimitCpu:          c.Int("limitCpu"),
		LimitMem:          c.Int("limitMem"),
		MinReplicas:       c.Int("minReplicas"),
		MaxReplicas:       c.Int("maxReplicas"),
		ServiceAccount:    c.String("serviceAccount"),
		IsNotUserCenter:   c.String("name") != "usercenter",
	})
	if err != nil {
		fmt.Printf("解析（%s）配置时发生错误：%s", c.String("name"), err.Error())
		return err
	}

	//fmt.Println(aurora.Green("Done."))
	return nil
}

// Category returns the category of the deployments.
func Category() string {
	return category
}

// Clean cleans the generated deployment files.
func Clean() error {
	return pathx.Clean(category)
}

// GenTemplates generates the deployment template files.
func GenTemplates(_ *cli.Context) error {
	return pathx.InitTemplates(category, map[string]string{
		deployTemplateFile: deploymentTemplate,
		jobTemplateFile:    jobTmeplate,
	})
}

// RevertTemplate reverts the given template file to the default value.
func RevertTemplate(name string) error {
	return pathx.CreateTemplate(category, name, deploymentTemplate)
}

// Update updates the template files to the templates built in current goctl.
func Update() error {
	err := Clean()
	if err != nil {
		return err
	}

	return pathx.InitTemplates(category, map[string]string{
		deployTemplateFile: deploymentTemplate,
		jobTemplateFile:    jobTmeplate,
	})
}
