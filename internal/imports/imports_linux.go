// Code generated by github.com/edwarnicke/imports-gen DO NOT EDIT.
package imports

import (
	_ "context"
	_ "crypto/tls"
	_ "fmt"
	_ "github.com/antonfisher/nested-logrus-formatter"
	_ "github.com/edwarnicke/debug"
	_ "github.com/edwarnicke/exechelper"
	_ "github.com/edwarnicke/genericsync"
	_ "github.com/edwarnicke/grpcfd"
	_ "github.com/go-ping/ping"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/google/uuid"
	_ "github.com/kelseyhightower/envconfig"
	_ "github.com/networkservicemesh/api/pkg/api/networkservice"
	_ "github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/cls"
	_ "github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/ipsec"
	_ "github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/kernel"
	_ "github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/memif"
	_ "github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/noop"
	_ "github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/vfio"
	_ "github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/vxlan"
	_ "github.com/networkservicemesh/api/pkg/api/networkservice/mechanisms/wireguard"
	_ "github.com/networkservicemesh/api/pkg/api/networkservice/payload"
	_ "github.com/networkservicemesh/api/pkg/api/registry"
	_ "github.com/networkservicemesh/govpp/binapi/acl"
	_ "github.com/networkservicemesh/govpp/binapi/acl_types"
	_ "github.com/networkservicemesh/govpp/binapi/af_packet"
	_ "github.com/networkservicemesh/govpp/binapi/af_xdp"
	_ "github.com/networkservicemesh/govpp/binapi/fib_types"
	_ "github.com/networkservicemesh/govpp/binapi/interface"
	_ "github.com/networkservicemesh/govpp/binapi/interface_types"
	_ "github.com/networkservicemesh/govpp/binapi/ip"
	_ "github.com/networkservicemesh/govpp/binapi/ip6_nd"
	_ "github.com/networkservicemesh/govpp/binapi/ip_neighbor"
	_ "github.com/networkservicemesh/govpp/binapi/ip_types"
	_ "github.com/networkservicemesh/sdk-k8s/pkg/tools/deviceplugin"
	_ "github.com/networkservicemesh/sdk-k8s/pkg/tools/podresources"
	_ "github.com/networkservicemesh/sdk-kernel/pkg/kernel/tools/nshandle"
	_ "github.com/networkservicemesh/sdk-sriov/pkg/networkservice/chains/forwarder"
	_ "github.com/networkservicemesh/sdk-sriov/pkg/networkservice/common/resourcepool"
	_ "github.com/networkservicemesh/sdk-sriov/pkg/sriov/config"
	_ "github.com/networkservicemesh/sdk-sriov/pkg/sriov/pci"
	_ "github.com/networkservicemesh/sdk-sriov/pkg/sriov/resource"
	_ "github.com/networkservicemesh/sdk-sriov/pkg/sriov/token"
	_ "github.com/networkservicemesh/sdk-sriov/pkg/tools/tokens"
	_ "github.com/networkservicemesh/sdk-sriov/pkg/tools/yamlhelper"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/networkservice/chains/forwarder"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/networkservice/connectioncontext"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/networkservice/mechanisms/ipsec"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/networkservice/mechanisms/memif"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/networkservice/mechanisms/vxlan"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/networkservice/mechanisms/wireguard"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/networkservice/metrics"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/networkservice/pinhole"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/networkservice/tag"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/networkservice/up"
	_ "github.com/networkservicemesh/sdk-vpp/pkg/tools/types"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/chains/client"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/chains/endpoint"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/common/authorize"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/common/cleanup"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/common/mechanisms"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/common/mechanisms/kernel"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/common/mechanisms/recvfd"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/common/mechanisms/sendfd"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/common/switchcase"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/core/chain"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/core/next"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/ipam/point2pointipam"
	_ "github.com/networkservicemesh/sdk/pkg/networkservice/utils/metadata"
	_ "github.com/networkservicemesh/sdk/pkg/registry/chains/client"
	_ "github.com/networkservicemesh/sdk/pkg/registry/common/authorize"
	_ "github.com/networkservicemesh/sdk/pkg/registry/common/begin"
	_ "github.com/networkservicemesh/sdk/pkg/registry/common/expire"
	_ "github.com/networkservicemesh/sdk/pkg/registry/common/memory"
	_ "github.com/networkservicemesh/sdk/pkg/registry/common/recvfd"
	_ "github.com/networkservicemesh/sdk/pkg/registry/common/sendfd"
	_ "github.com/networkservicemesh/sdk/pkg/registry/core/adapters"
	_ "github.com/networkservicemesh/sdk/pkg/registry/core/chain"
	_ "github.com/networkservicemesh/sdk/pkg/tools/grpcutils"
	_ "github.com/networkservicemesh/sdk/pkg/tools/log"
	_ "github.com/networkservicemesh/sdk/pkg/tools/log/logruslogger"
	_ "github.com/networkservicemesh/sdk/pkg/tools/monitorconnection/authorize"
	_ "github.com/networkservicemesh/sdk/pkg/tools/opentelemetry"
	_ "github.com/networkservicemesh/sdk/pkg/tools/pprofutils"
	_ "github.com/networkservicemesh/sdk/pkg/tools/prometheus"
	_ "github.com/networkservicemesh/sdk/pkg/tools/spiffejwt"
	_ "github.com/networkservicemesh/sdk/pkg/tools/spire"
	_ "github.com/networkservicemesh/sdk/pkg/tools/token"
	_ "github.com/networkservicemesh/sdk/pkg/tools/tracing"
	_ "github.com/networkservicemesh/vpphelper"
	_ "github.com/networkservicemesh/vpphelper/extendtimeout"
	_ "github.com/pkg/errors"
	_ "github.com/safchain/ethtool"
	_ "github.com/sirupsen/logrus"
	_ "github.com/spiffe/go-spiffe/v2/bundle/x509bundle"
	_ "github.com/spiffe/go-spiffe/v2/spiffeid"
	_ "github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	_ "github.com/spiffe/go-spiffe/v2/svid/x509svid"
	_ "github.com/spiffe/go-spiffe/v2/workloadapi"
	_ "github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/suite"
	_ "github.com/thanhpk/randstr"
	_ "github.com/vishvananda/netlink"
	_ "github.com/vishvananda/netns"
	_ "go.fd.io/govpp/adapter/statsclient"
	_ "go.fd.io/govpp/api"
	_ "go.fd.io/govpp/binapi/vlib"
	_ "go.fd.io/govpp/binapi/vpe"
	_ "go.fd.io/govpp/core"
	_ "golang.org/x/text/cases"
	_ "golang.org/x/text/language"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/health/grpc_health_v1"
	_ "gopkg.in/yaml.v3"
	_ "net"
	_ "net/url"
	_ "os"
	_ "os/signal"
	_ "path"
	_ "path/filepath"
	_ "runtime"
	_ "strconv"
	_ "strings"
	_ "syscall"
	_ "testing"
	_ "time"
)
