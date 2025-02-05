package instantquery

import (
	"testing"

	"github.com/cortexproject/cortex/pkg/querier/tripperware"
	"github.com/cortexproject/cortex/pkg/querier/tripperware/queryrange"
)

func Test_shardQuery(t *testing.T) {
	tripperware.TestQueryShardQuery(t, InstantQueryCodec, queryrange.ShardedPrometheusCodec)
}
