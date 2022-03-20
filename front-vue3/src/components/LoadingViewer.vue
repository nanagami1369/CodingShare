<template>
  <div class="loading-viewer">
    <div class="loading-viewer-title">読込中……</div>
    <p>大きなデータを読み込んでいます</p>
    <p>しばらくお待ち下さい</p>
  </div>
</template>

<style scoped>
.loading-viewer {
  display: flex;
  flex-direction: column;
  background-color: white;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 6px 40px;
  z-index: 1000;
  width: 80vh;
  border-radius: 10px;
  border: solid 5px;
}

.loading-viewer-title {
  text-align: left;
  padding-bottom: 15px;
}
.loading-viewer p {
  text-align: left;
}
</style>
