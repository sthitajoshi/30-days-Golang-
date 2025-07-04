// package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"path/filepath"

// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	v1 "k8s.io/client-go/applyconfigurations/core/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// )

// func main() {
// 	// fetch kubeconfig file
// 	home, _ := os.UserHomeDir()
// 	kubeconfigPath := filepath.Join(home,".kube/config")

// 	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
// 	if err!= nil {
// 		panic(err.Error())
// 	}
// 	// create the clientset
// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	_ = clientset // use clientset as needed

// 	// //read existing pods (have a knowlage of struct and pointer)
// 	// pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }
// 	// for i, pod := range pods.Items {
// 	// 	fmt.Printf("name of the %dth pod: %s\n", i, pod.Name)
// 	// }

// 	// cereate a new pods

// 	podDefinition := v1.Pod("demo-k8s-", "default").
// 		WithSpec(
// 			v1.PodSpec().
// 				WithContainers(
// 					v1.Container().
// 						WithName("nginx-container").
// 						WithImage("nginx:latest"),
// 				),
// 		)

// 	// Create the Pod in the cluster
// 	createdPod, err := clientset.CoreV1().Pods("default").Apply(
// 		context.Background(),
// 		podDefinition,
// 		metav1.ApplyOptions{
// 			FieldManager: "example-controller",
// 		},
// 	)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Printf("Created pod %s\n", createdPod.Name)
// }

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Initialize Kubernetes client
func initKubernetesClient() (*kubernetes.Clientset, error) {
	home, _ := os.UserHomeDir()
	kubeconfigPath := filepath.Join(home, ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

// List existing pods
func listPods(clientset *kubernetes.Clientset) {
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Existing pods:")
	for i, pod := range pods.Items {
		fmt.Printf("  %d. %s (Status: %s)\n", i+1, pod.Name, pod.Status.Phase)
	}
}

// Create a new pod
func createPod(clientset *kubernetes.Clientset) string {
	podName := fmt.Sprintf("demo-k8s-pod-%d", time.Now().Unix())

	podDefinition := v1.Pod(podName, "default").
		WithSpec(
			v1.PodSpec().
				WithContainers(
					v1.Container().
						WithName("nginx-container").
						WithImage("nginx:latest"),
				),
		)

	createdPod, err := clientset.CoreV1().Pods("default").Apply(
		context.Background(),
		podDefinition,
		metav1.ApplyOptions{
			FieldManager: "example-controller",
		},
	)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("✅ Created pod: %s\n", createdPod.Name)
	return createdPod.Name
}

// Update a pod (change image)
func updatePod(clientset *kubernetes.Clientset, podName string) {
	fmt.Printf("🔄 Updating pod: %s\n", podName)

	updatedPodDefinition := v1.Pod(podName, "default").
		WithSpec(
			v1.PodSpec().
				WithContainers(
					v1.Container().
						WithName("nginx-container").
						WithImage("nginx:1.21"), // Updated image version
				),
		)

	updatedPod, err := clientset.CoreV1().Pods("default").Apply(
		context.Background(),
		updatedPodDefinition,
		metav1.ApplyOptions{
			FieldManager: "example-controller",
		},
	)
	if err != nil {
		fmt.Printf("❌ Error updating pod %s: %v\n", podName, err)
		return
	}

	fmt.Printf("✅ Updated pod: %s\n", updatedPod.Name)
}

// Delete a pod
func deletePod(clientset *kubernetes.Clientset, podName string) {
	fmt.Printf("🗑️ Deleting pod: %s\n", podName)

	err := clientset.CoreV1().Pods("default").Delete(
		context.Background(),
		podName,
		metav1.DeleteOptions{},
	)
	if err != nil {
		fmt.Printf("❌ Error deleting pod %s: %v\n", podName, err)
		return
	}

	fmt.Printf("✅ Successfully deleted pod: %s\n", podName)
}

func main() {
	// Initialize Kubernetes client
	clientset, err := initKubernetesClient()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("=== Initial Pod List ===")
	listPods(clientset)

	// Create a new pod
	fmt.Println("\n=== Creating Pod ===")
	podName := createPod(clientset)

	// Wait for pod to be created
	fmt.Println("\nWaiting 3 seconds for pod to initialize...")
	time.Sleep(3 * time.Second)

	fmt.Println("\n=== Pod List After Creation ===")
	listPods(clientset)

	// Update the pod
	fmt.Println("\n=== Updating Pod ===")
	updatePod(clientset, podName)

	// Wait for update to process
	fmt.Println("\nWaiting 3 seconds for update to process...")
	time.Sleep(3 * time.Second)

	fmt.Println("\n=== Pod List After Update ===")
	listPods(clientset)

	// Delete the pod
	fmt.Println("\n=== Deleting Pod ===")
	deletePod(clientset, podName)

	// Wait for deletion to complete
	fmt.Println("\nWaiting 3 seconds for deletion to complete...")
	time.Sleep(3 * time.Second)

	fmt.Println("\n=== Final Pod List ===")
	listPods(clientset)
}
