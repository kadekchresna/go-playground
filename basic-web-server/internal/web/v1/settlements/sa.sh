# Step 1: Create a ServiceAccount
kubectl create serviceaccount go-job-sa

# Step 2: Create a ClusterRoleBinding
kubectl create clusterrolebinding my-service-account-binding --clusterrole=admin --serviceaccount=default:go-job-sa

# Step 3: Get the Secret name
SECRET_NAME=$(kubectl get serviceaccount go-job-sa -o jsonpath='{.secrets[0].name}')

# Step 4: Extract the token
TOKEN=$(kubectl get secret $SECRET_NAME -o jsonpath='{.data.token}' | base64 --decode)

echo "Your static token is: $TOKEN"