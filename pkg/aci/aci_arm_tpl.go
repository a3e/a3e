package aci

type aciARMTpl struct {
	APIVersion string          `json:"apiVersion"`
	Identity   aciARMTpl_sub2  `json:"identity"`
	Location   string          `json:"location"`
	Name       string          `json:"name"`
	Properties aciARMTpl_sub23 `json:"properties"`
	Tags       aciARMTpl_sub1  `json:"tags"`
	Type       string          `json:"type"`
}

type aciARMTpl_sub22 struct {
	AzureFile aciARMTpl_sub20 `json:"azureFile"`
	EmptyDir  aciARMTpl_sub1  `json:"emptyDir"`
	GitRepo   aciARMTpl_sub21 `json:"gitRepo"`
	Name      string          `json:"name"`
	Secret    aciARMTpl_sub1  `json:"secret"`
}

type aciARMTpl_sub9 struct {
	CPU        string         `json:"cpu"`
	Gpu        aciARMTpl_sub8 `json:"gpu"`
	MemoryInGB string         `json:"memoryInGB"`
}

type aciARMTpl_sub12 struct {
	Command              []string          `json:"command"`
	EnvironmentVariables []aciARMTpl_sub3  `json:"environmentVariables"`
	Image                string            `json:"image"`
	LivenessProbe        aciARMTpl_sub6    `json:"livenessProbe"`
	Ports                []aciARMTpl_sub7  `json:"ports"`
	ReadinessProbe       aciARMTpl_sub6    `json:"readinessProbe"`
	Resources            aciARMTpl_sub10   `json:"resources"`
	VolumeMounts         []aciARMTpl_sub11 `json:"volumeMounts"`
}

type aciARMTpl_sub4 struct {
	Command []string `json:"command"`
}

type aciARMTpl_sub23 struct {
	Containers               []aciARMTpl_sub13 `json:"containers"`
	Diagnostics              aciARMTpl_sub15   `json:"diagnostics"`
	DNSConfig                aciARMTpl_sub16   `json:"dnsConfig"`
	ImageRegistryCredentials []aciARMTpl_sub17 `json:"imageRegistryCredentials"`
	IPAddress                aciARMTpl_sub18   `json:"ipAddress"`
	NetworkProfile           aciARMTpl_sub19   `json:"networkProfile"`
	OsType                   string            `json:"osType"`
	RestartPolicy            string            `json:"restartPolicy"`
	Volumes                  []aciARMTpl_sub22 `json:"volumes"`
}

type aciARMTpl_sub8 struct {
	Count string `json:"count"`
	Sku   string `json:"sku"`
}

type aciARMTpl_sub18 struct {
	DNSNameLabel string           `json:"dnsNameLabel"`
	IP           string           `json:"ip"`
	Ports        []aciARMTpl_sub7 `json:"ports"`
	Type         string           `json:"type"`
}

type aciARMTpl_sub21 struct {
	Directory  string `json:"directory"`
	Repository string `json:"repository"`
	Revision   string `json:"revision"`
}

type aciARMTpl_sub6 struct {
	Exec                aciARMTpl_sub4 `json:"exec"`
	FailureThreshold    string         `json:"failureThreshold"`
	HTTPGet             aciARMTpl_sub5 `json:"httpGet"`
	InitialDelaySeconds string         `json:"initialDelaySeconds"`
	PeriodSeconds       string         `json:"periodSeconds"`
	SuccessThreshold    string         `json:"successThreshold"`
	TimeoutSeconds      string         `json:"timeoutSeconds"`
}

type aciARMTpl_sub19 struct {
	ID string `json:"id"`
}

type aciARMTpl_sub10 struct {
	Limits   aciARMTpl_sub9 `json:"limits"`
	Requests aciARMTpl_sub9 `json:"requests"`
}

type aciARMTpl_sub15 struct {
	LogAnalytics aciARMTpl_sub14 `json:"logAnalytics"`
}

type aciARMTpl_sub14 struct {
	LogType      string         `json:"logType"`
	Metadata     aciARMTpl_sub1 `json:"metadata"`
	WorkspaceID  string         `json:"workspaceId"`
	WorkspaceKey string         `json:"workspaceKey"`
}

type aciARMTpl_sub11 struct {
	MountPath string `json:"mountPath"`
	Name      string `json:"name"`
	ReadOnly  string `json:"readOnly"`
}

type aciARMTpl_sub13 struct {
	Name       string          `json:"name"`
	Properties aciARMTpl_sub12 `json:"properties"`
}

type aciARMTpl_sub3 struct {
	Name        string `json:"name"`
	SecureValue string `json:"secureValue"`
	Value       string `json:"value"`
}

type aciARMTpl_sub16 struct {
	NameServers   []string `json:"nameServers"`
	Options       string   `json:"options"`
	SearchDomains string   `json:"searchDomains"`
}

type aciARMTpl_sub17 struct {
	Password string `json:"password"`
	Server   string `json:"server"`
	Username string `json:"username"`
}

type aciARMTpl_sub5 struct {
	Path   string `json:"path"`
	Port   string `json:"port"`
	Scheme string `json:"scheme"`
}

type aciARMTpl_sub7 struct {
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
}

type aciARMTpl_sub20 struct {
	ReadOnly           string `json:"readOnly"`
	ShareName          string `json:"shareName"`
	StorageAccountKey  string `json:"storageAccountKey"`
	StorageAccountName string `json:"storageAccountName"`
}

type aciARMTpl_sub2 struct {
	Type                   string         `json:"type"`
	UserAssignedIdentities aciARMTpl_sub1 `json:"userAssignedIdentities"`
}

type aciARMTpl_sub1 struct{}
