<template>
  <div class="not-found-page">
    <div>
      <p>Not found</p>
      <p>そのURLは存在しません</p>
      <p>
        Top Page→ <router-link to="/">{{ topPageUrl }}</router-link>
      </p>
    </div>
  </div>
</template>
<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'NotFoundPage',
  computed: {
    topPageUrl: function (): string {
      return `${location.host}${process.env.BASE_URL}#/`
    },
  },
})
</script>
<style scoped>
.not-found-page {
  font-size: 3em;
  display: flex;
  justify-content: center;
  align-items: center;
}
.not-found-page pre {
  text-align: left;
}
</style>
