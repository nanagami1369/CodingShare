<template>
  <section class="thumbnail" @click="openPlayer">
    <div class="thumbnail-video-info">
      <div>
        <span class="thumbnail-video-info-label">タイトル</span>
        <a
          ><cite class="thumbnail-video-info-value">
            {{ video.header.title }}
          </cite>
        </a>
      </div>
      <div>
        <span class="thumbnail-video-info-label">作者</span>
        <span @click.stop>
          <router-link
            :to="{ name: 'UserPage', params: { id: video.header.name } }"
          >
            <cite class="thumbnail-video-info-value">
              {{ video.header.name }}
            </cite>
          </router-link>
        </span>
      </div>
      <div>
        <span class="thumbnail-video-info-label">言語</span>
        <span class="thumbnail-video-info-value">
          {{ video.header.language.name }}
        </span>
      </div>
      <div>
        <span class="thumbnail-video-info-label"> 投稿日</span>
        <span class="thumbnail-video-info-value"> {{ uploadDate }}</span>
      </div>
      <div>
        <span class="thumbnail-video-info-label">時間</span>
        <sapn class="thumbnail-video-info-value"> {{ recordingTime }}</sapn>
      </div>
    </div>
    <img src="@/assets/c_thumbnail.png" class="thumbnail-image" />
  </section>
</template>

<script lang="ts">
import { Video } from '@/models/Video'
import { formatRecordingTime } from '@/util'
import { format } from 'date-fns'
import Vue, { PropType } from 'vue'
export default Vue.extend({
  name: 'Thumbnail',
  props: {
    video: {
      type: Object as PropType<Video>,
      required: true,
    },
  },
  computed: {
    uploadDate: function (): string {
      return format(this.video.header.uploadTime, 'yyyy/MM/dd')
    },
    recordingTime: function (): string {
      return formatRecordingTime(this.video.header.recordingTime)
    },
  },
  methods: {
    openPlayer: function (): void {
      this.$router.push({
        name: 'Player',
        params: { id: this.video.header.videoId + '' },
      })
    },
  },
})
</script>

<style scoped>
.thumbnail {
  border: solid 2px #202020;
  width: 550px;
  margin: 10px 20px;
  padding: 5px 10px;
  display: flex;
  justify-content: space-evenly;
  align-items: center;
  cursor: pointer;
}
.thumbnail-video-info {
  display: block;
  text-align: left;
  flex: 1 1 auto;
}

.thumbnail-video-info-label {
  display: inline-block;
  text-align: right;
  width: 5em;
}
.thumbnail-video-info-label::after {
  content: '｜';
}
.thumbnail-video-info-value {
  overflow-wrap: break-word;
  word-break: break-all;
}
img.thumbnail-image {
  size: 80%;
  flex: 0 0 auto;
  height: 110px;
  width: 210px;
}
</style>
