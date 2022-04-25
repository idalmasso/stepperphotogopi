<template>
  <div class="ended-process-viewer">
    <h2>Processo {{ processName }}</h2>
    <vue-three-sixty
      v-if="savewebp"
      :amount="numImages"
      :imagePath="imageDir"
      fileName="{index}.webp"
    />
    <vue-three-sixty
      v-else
      :amount="numImages"
      :imagePath="imageDir"
      fileName="{index}.jpg"
    />
  </div>
</template>

<script>
export default {
  name: 'EndedProcessViewer',
  props: {
    processName: String,
    numImages: Number,
  },
  data() {
    return {
      savewebp: false,
    }
  },
  computed: {
    imageDir() {
      //console.log(this.processName)
      return '/process-images/' + this.processName + '/'
    },
  },
  mounted() {
    fetch('/api/configuration')
      .then((data) => {
        return data.json()
      })
      .then((config) => {
        this.savewebp = config.server.saveAsWebP
      })
  },
}
</script>

<style>
@import '../assets/css/all.css';
</style>
