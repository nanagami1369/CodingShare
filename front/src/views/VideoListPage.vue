<template>
  <div>
    <h1>動画一覧</h1>
    <ul>
      <li v-for="videoName in listValue" :key="videoName">
        <router-link :to="{ name: 'Player', query: { src: videoName } }">{{
          videoName
        }}</router-link>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

type DataType = {
  listValue: string[]
}

export default Vue.extend({
  name: 'VideoListPage',
  data(): DataType {
    return {
      listValue: ['sample.json'],
    }
  },
})
</script>
