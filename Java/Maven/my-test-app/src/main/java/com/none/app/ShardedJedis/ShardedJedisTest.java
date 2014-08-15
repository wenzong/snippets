package com.none.app.ShardedJedis;

import java.net.MalformedURLException;
import java.net.URISyntaxException;
import java.net.URL;
import java.util.ArrayList;
import java.util.List;

import org.apache.commons.pool2.impl.GenericObjectPoolConfig;

import redis.clients.jedis.JedisShardInfo;
import redis.clients.jedis.ShardedJedis;
import redis.clients.jedis.ShardedJedisPool;

class ShardedJedisTest{
    public static void main(String[] args) throws URISyntaxException,
            MalformedURLException {
        List<JedisShardInfo> shards = new ArrayList<JedisShardInfo>();

        JedisShardInfo si = new JedisShardInfo("10.10.3.101", 6465);
        shards.add(si);
        si = new JedisShardInfo("10.10.3.123", 6465);
        shards.add(si);

        ShardedJedisPool pool = new ShardedJedisPool(
                new GenericObjectPoolConfig(), shards);

        ShardedJedis jedis = pool.getResource();
        jedis.set("a", "got u");
        System.out.println(jedis.get("a"));
        jedis.del("a");
        pool.returnResource(jedis);
        
        pool.destroy();
    }
}
