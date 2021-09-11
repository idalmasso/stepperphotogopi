<template>
  <button @click="fetchImage">Request image</button>
  <img v-if="imageObjectURL != ''" :src="imageObjectURL" />
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
      fetch(this.imageUrl)
        .then((response) => response.blob())
        .then((imageBlob) => {
          // Then create a local URL for that image and print it
          this.imageObjectURL = URL.createObjectURL(imageBlob)
          console.log(this.imageObjectURL)
        })
    },
  },
  computed: {
    imageUrl() {
      return '/api/get-snapshot'
    },
  },
}
</script>

<style></style>
