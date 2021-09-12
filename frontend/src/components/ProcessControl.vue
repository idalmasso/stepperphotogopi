<template>
  <div>
    <h1>Photo processing</h1>
    <div v-if="error">{{ error }}</div>
    <form v-if="motorStatus != 'working'">
      <button type="button" @click="submit">START</button>
    </form>
    <button v-if="motorStatus == 'working'" @click="stopProcess">
      Stop process
    </button>
    <h4>Motor status: {{ motorStatus }}</h4>
  </div>
</template>

<script>
export default {
  name: 'ProcessControl',
  data() {
    return {
      degrees: 0,
      error: '',
      motorStatus: '',
    }
  },
  methods: {
    stopProcess() {
      fetch('/api/stop-process', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
      })
        .then(async (a) => {
          if (!a.ok) {
            var err = await a.text()
            throw Error(err)
          }
        })
        .catch((error) => {
          this.error = error
        })
    },
    submit() {
      fetch('/api/start-process', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
      })
        .then(async (a) => {
          if (!a.ok) {
            var err = await a.text()
            throw Error(err)
          }
        })
        .catch((error) => {
          this.error = error
        })
    },
    requestStatus() {
      fetch('/api/machine-status')
        .then(async (a) => {
          if (!a.ok) {
            var err = await a.text()
            throw Error(err)
          }
          return a.json()
        })
        .then((data) => {
          this.motorStatus = data['value']
        })
        .catch((error) => {
          this.error = error
        })
    },
  },
  mounted() {
    setInterval(this.requestStatus, 1000)
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
