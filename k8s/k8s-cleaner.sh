docker ps -a
service kubelet stop
service kube-proxy stop
service etcd stop
service flanneld stop
#service calico-node status
service calico-node stop
docker kill `docker ps -q`
docker rm `docker ps -aq`
service docker stop
ip link del flannel.1
killall -9 kubemark
# ip route | grep 172.16. | awk '{print $0}' | while read m; do ip route del $m; done
for m in `df -h | grep ^tmpfs | grep "/var/lib/kubelet/pods" | awk '{print $6}'`; do umount $m; done
for m in `mount | grep kubelet | awk {'print $3'}`; do umount $m; done
rm -rf /etc/kubernetes
rm -rf /var/lib/docker
rm -rf /var/lib/kubelet
rm -rf /var/etcd/data/*
rm -rf /var/run/kubernetes
rm -rf /var/run/flannel
rm -rf /var/run/docker
rm -f /etc/network/if-up.d/vip
rm -f /etc/init/calico-node.conf
rm -f /etc/cni/net.d/10-calico.conf
rm -f /etc/init/kubelet.conf
rm -f /etc/init/kube-proxy.conf
grep -v -i flannel /etc/default/docker > /tmp/docker
mv /tmp/docker /etc/default/docker
killall -9 ucarp


# list status of all the NotReady nodes:
for m in `kubectl get nodes | grep NotReady | awk {'print $1'}`; do kubectl get node $m -owide; done

# remove all the NotReady nodes:
for m in `kubectl get nodes | grep NotReady | awk {'print $1'}`; do kubectl delete node $m; done
