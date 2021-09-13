<template>
  <div class="test-page">
    <h1>Test page</h1>
    <MotorControl />
    <CameraView />
  </div>
</template>

<script>
// @ is an alias to /src
import MotorControl from '@/components/MotorControl.vue'
import CameraView from '@/components/CameraView.vue'

export default {
  name: 'TestView',
  components: {
    MotorControl,
    CameraView,
  },
}
</script>
