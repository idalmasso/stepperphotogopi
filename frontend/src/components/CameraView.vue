<template>
  <img class="camera-image" v-if="imageObjectURL != ''" :src="imageObjectURL" />
  <button @click="swapInterval">{{ swapIntervalText }}</button>
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
}
</script>

<style>
.camera-image {
  max-width: 50%;
}
</style>
