<template>
  <header class="header">
    <button class="header-button" @click="changeHomePage">
      <FontAwesomeIcon icon="home" />
    </button>
  </header>
</template>
<script lang="ts">
import Vue from 'vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'
import { faHome } from '@fortawesome/free-solid-svg-icons'

library.add(faHome)
export default Vue.extend({
  name: 'HomePage',
  components: {
    FontAwesomeIcon,
  },
  methods: {
    changeHomePage: function (): void {
      this.$router.push({ path: '/' })
    },
  },
})
</script>

<style scoped>
.header {
  background: #202020;
  color: #eeeeee;
  display: flex;
  padding: 0px 15px;
}
.header-button {
  cursor: pointer;
  color: #eeeeee;
  height: 100%;
  font-size: 1.5em;
  flex: 0 0 auto;
  margin: 0px 5px;
  background-color: transparent;
  border: none;
}
</style>
