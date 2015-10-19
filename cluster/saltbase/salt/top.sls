base:
  '*':
    - base
    - debian-auto-upgrades
    - salt-helpers
{% if grains.get('cloud') == 'aws' %}
    - ntp
{% endif %}

  'roles:kubernetes-pool':
    - match: grain
    - docker
{% if grains['cloud'] is defined and grains['cloud'] == 'azure' %}
    - openvpn-client
{% endif %}
    - helpers
    - cadvisor
    - kube-client-tools
{% if pillar.get('network_provider', '').lower() == 'opencontrail' %}
    - opencontrail-kubelet-plugin
{% endif %}
    - kubelet
{% if pillar.get('network_provider', '').lower() == 'opencontrail' %}
    - opencontrail-networking-minion
{% else %}
    - kube-proxy
{% endif %}
{% if pillar.get('enable_node_logging', '').lower() == 'true' and pillar['logging_destination'] is defined %}
  {% if pillar['logging_destination'] == 'elasticsearch' %}
    - fluentd-es
  {% elif pillar['logging_destination'] == 'gcp' %}
    - fluentd-gcp
  {% endif %}
{% endif %}
{% if pillar.get('enable_cluster_registry', '').lower() == 'true' %}
    - kube-registry-proxy
{% endif %}
    - logrotate
{% if grains['cloud'] is defined and grains.cloud == 'gce' %}
    - supervisor
{% else %}
    - monit
{% endif %}

  'roles:kubernetes-master':
    - match: grain
    - generate-cert
    - etcd
{% if grains['cloud'] is defined and grains['cloud'] in [ 'vagrant', 'gce', 'aws' ] %}
    - docker
    - kubelet
{% endif %}
{% if grains['cloud'] is defined and grains.cloud == 'gce' %}
    - supervisor
{% else %}
    - monit
{% endif %}
    - kube-apiserver
    - kube-controller-manager
    - kube-scheduler
{% if grains['cloud'] is defined and not grains.cloud in [ 'aws', 'gce', 'vagrant' ] %}
    - nginx
{% endif %}
    - cadvisor
    - kube-client-tools
    - kube-master-addons
    - kube-admission-controls
{% if pillar.get('enable_node_logging', '').lower() == 'true' and pillar['logging_destination'] is defined %}
  {% if pillar['logging_destination'] == 'elasticsearch' %}
    - fluentd-es
  {% elif pillar['logging_destination'] == 'gcp' %}
    - fluentd-gcp
  {% endif %}
{% endif %}
{% if grains['cloud'] is defined and grains['cloud'] != 'vagrant' %}
    - logrotate
{% endif %}
    - kube-addons
{% if grains['cloud'] is defined and grains['cloud'] == 'azure' %}
    - openvpn
{% endif %}
{% if pillar.get('network_provider', '').lower() == 'opencontrail' %}
    - opencontrail-networking-master
{% endif %}

  'roles:kubernetes-pool-vsphere':
    - match: grain
    - static-routes

  'roles:kubernetes-network-provider-gateway':
    - match: grain
{% if pillar.get('network_provider', '').lower() == 'opencontrail' %}
    - opencontrail-networking-gateway
{% endif %}
