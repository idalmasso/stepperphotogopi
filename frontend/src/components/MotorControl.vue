<template>
  <div>
    <div :v-if="error">{{ error }}</div>
    <form>
      <label for="degreerequest">Degrees</label>|
      <input id="degreerequest" type="text" v-model="degrees" />
      <button type="button" @click="submit">Submit</button>
    </form>
    <button v-if="motorStatus == 'working'" @click="stopMotor">
      Stop motor
    </button>
    <h4>Motor status: {{ motorStatus }}</h4>
  </div>
</template>

<script>
export default {
  name: 'MotorControl',
  data() {
    return {
      degrees: 0,
      error: '',
      motorStatus: '',
    }
  },
  methods: {
    stopMotor() {
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
      fetch('/api/move-motor', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ degrees: this.degrees }),
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
