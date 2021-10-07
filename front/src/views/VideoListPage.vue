<template>
  <div>
    <h1>動画一覧</h1>
    <ul>
      <li>
        <router-link :to="{ name: 'Player', query: { src: 'sample.json' } }"
          >サンプル動画</router-link
        >
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  name: 'VideoListPage',
})
</script>
