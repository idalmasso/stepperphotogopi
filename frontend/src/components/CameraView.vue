<template>
<div class="camera-view">
  <h2>Camera test</h2>
  <img class="camera-image" v-if="imageObjectURL != ''" :src="imageObjectURL" />
  <button @click="swapInterval">{{ swapIntervalText }}</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      imageObjectURL: '',
      interval: null,
    }
  },
  methods: {
    fetchImage() {
      console.log(this.imageObjectURL)
      fetch(this.imageUrl)
        .then((response) => response.blob())
        .then((imageBlob) => {
          // Then create a local URL for that image and print it
          this.imageObjectURL = URL.createObjectURL(imageBlob)
          console.log(this.imageObjectURL)
        })
        .catch((e) => {
          console.log(e)
        })
    },
    swapInterval() {
      if (this.interval == null) {
        this.interval = setInterval(this.fetchImage, 3000)
      } else {
        clearInterval(this.interval)
        this.interval = null
      }
    },
  },
  computed: {
    imageUrl() {
      return '/api/get-snapshot'
    },
    swapIntervalText() {
      if (this.interval == null) {
        return 'Raccogli immagini'
      }
      return 'stop'
    },
  },
  beforeUnmount() {
    if (this.interval != null) {
      clearInterval(this.interval)
    }
  },
}
</script>

<style>
.camera-image {
  max-width: 50%;
}
.camera-view {
  border-style: none solid solid solid;
  border-width: 1pt;
  padding-bottom: 10pt;
}
</style>
