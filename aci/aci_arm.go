package main

type ACI struct {
	APIVersion string `json:"apiVersion"`
	Identity   struct {
		Type                   string   `json:"type"`
		UserAssignedIdentities struct{} `json:"userAssignedIdentities"`
	} `json:"identity"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	Properties struct {
		Containers []struct {
			Name       string `json:"name"`
			Properties struct {
				Command              []string `json:"command"`
				EnvironmentVariables []struct {
					Name        string `json:"name"`
					SecureValue string `json:"secureValue"`
					Value       string `json:"value"`
				} `json:"environmentVariables"`
				Image         string `json:"image"`
				LivenessProbe struct {
					Exec struct {
						Command []string `json:"command"`
					} `json:"exec"`
					FailureThreshold string `json:"failureThreshold"`
					HTTPGet          struct {
						Path   string `json:"path"`
						Port   string `json:"port"`
						Scheme string `json:"scheme"`
					} `json:"httpGet"`
					InitialDelaySeconds string `json:"initialDelaySeconds"`
					PeriodSeconds       string `json:"periodSeconds"`
					SuccessThreshold    string `json:"successThreshold"`
					TimeoutSeconds      string `json:"timeoutSeconds"`
				} `json:"livenessProbe"`
				Ports []struct {
					Port     string `json:"port"`
					Protocol string `json:"protocol"`
				} `json:"ports"`
				ReadinessProbe struct {
					Exec struct {
						Command []string `json:"command"`
					} `json:"exec"`
					FailureThreshold string `json:"failureThreshold"`
					HTTPGet          struct {
						Path   string `json:"path"`
						Port   string `json:"port"`
						Scheme string `json:"scheme"`
					} `json:"httpGet"`
					InitialDelaySeconds string `json:"initialDelaySeconds"`
					PeriodSeconds       string `json:"periodSeconds"`
					SuccessThreshold    string `json:"successThreshold"`
					TimeoutSeconds      string `json:"timeoutSeconds"`
				} `json:"readinessProbe"`
				Resources struct {
					Limits struct {
						CPU string `json:"cpu"`
						Gpu struct {
							Count string `json:"count"`
							Sku   string `json:"sku"`
						} `json:"gpu"`
						MemoryInGB string `json:"memoryInGB"`
					} `json:"limits"`
					Requests struct {
						CPU string `json:"cpu"`
						Gpu struct {
							Count string `json:"count"`
							Sku   string `json:"sku"`
						} `json:"gpu"`
						MemoryInGB string `json:"memoryInGB"`
					} `json:"requests"`
				} `json:"resources"`
				VolumeMounts []struct {
					MountPath string `json:"mountPath"`
					Name      string `json:"name"`
					ReadOnly  string `json:"readOnly"`
				} `json:"volumeMounts"`
			} `json:"properties"`
		} `json:"containers"`
		Diagnostics struct {
			LogAnalytics struct {
				LogType      string   `json:"logType"`
				Metadata     struct{} `json:"metadata"`
				WorkspaceID  string   `json:"workspaceId"`
				WorkspaceKey string   `json:"workspaceKey"`
			} `json:"logAnalytics"`
		} `json:"diagnostics"`
		DNSConfig struct {
			NameServers   []string `json:"nameServers"`
			Options       string   `json:"options"`
			SearchDomains string   `json:"searchDomains"`
		} `json:"dnsConfig"`
		ImageRegistryCredentials []struct {
			Password string `json:"password"`
			Server   string `json:"server"`
			Username string `json:"username"`
		} `json:"imageRegistryCredentials"`
		IPAddress struct {
			DNSNameLabel string `json:"dnsNameLabel"`
			IP           string `json:"ip"`
			Ports        []struct {
				Port     string `json:"port"`
				Protocol string `json:"protocol"`
			} `json:"ports"`
			Type string `json:"type"`
		} `json:"ipAddress"`
		NetworkProfile struct {
			ID string `json:"id"`
		} `json:"networkProfile"`
		OsType        string `json:"osType"`
		RestartPolicy string `json:"restartPolicy"`
		Volumes       []struct {
			AzureFile struct {
				ReadOnly           string `json:"readOnly"`
				ShareName          string `json:"shareName"`
				StorageAccountKey  string `json:"storageAccountKey"`
				StorageAccountName string `json:"storageAccountName"`
			} `json:"azureFile"`
			EmptyDir struct{} `json:"emptyDir"`
			GitRepo  struct {
				Directory  string `json:"directory"`
				Repository string `json:"repository"`
				Revision   string `json:"revision"`
			} `json:"gitRepo"`
			Name   string   `json:"name"`
			Secret struct{} `json:"secret"`
		} `json:"volumes"`
	} `json:"properties"`
	Tags struct{} `json:"tags"`
	Type string   `json:"type"`
}
