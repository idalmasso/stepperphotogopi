<template>
  <img class="camera-image" v-if="imageObjectURL != ''" :src="imageObjectURL" />
</template>

<script>
export default {
  data() {
    return {
      imageObjectURL: '',
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
  },
  computed: {
    imageUrl() {
      return '/api/get-snapshot'
    },
  },
  mounted() {
    setInterval(this.fetchImage, 3000)
  },
}
</script>

<style>
.camera-image {
  max-width: 50%;
}
</style>
