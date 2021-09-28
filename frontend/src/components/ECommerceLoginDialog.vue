<template>
  <transition name="modal">
    <div class="modal-mask">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header">
              <h2>Login in ecommerce</h2>
              <span class="error" v-if="message != ''">{{ message }}</span>
            </slot>
          </div>
          <div class="modal-body">
            <slot name="body">
              <table>
                <tr>
                  <td>
                    <label for="email">email</label>
                  </td>
                  <td><input type="text" id="email" v-model="email" /></td>
                </tr>
                <tr>
                  <td><label for="password">password</label></td>
                  <td>
                    <input type="password" id="password" v-model="password" />
                  </td>
                </tr>
              </table>
            </slot>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              <button class="modal-default-button" @click="cancelClicked">
                CANCEL
              </button>
              <button class="modal-default-button" @click="okClicked">
                OK
              </button>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: 'ECommerceLoginDialog',
  props: ['authUrl', 'defaultUsername', 'defaultPassword'],
  data() {
    return {
      message: '',
      email: '',
      password: '',
    }
  },
  methods: {
    okClicked() {
      fetch(this.authUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email: this.email, password: this.password }),
      })
        .then((a) => {
          return a.json()
        })
        .then((data) => {
          if (data.status === 'ok') {
            this.$emit('logged', data.token)
          } else {
            this.message = 'Cannot login: ' + data.status
          }
        })
    },
    cancelClicked() {
      this.$emit('close')
    },
  },
  mounted() {
    this.email = this.defaultUsername
    this.password = this.defaultPassword
  },
}
</script>

<style scoped>
.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: table;
  transition: opacity 0.3s ease;
}

.modal-wrapper {
  display: table-cell;
  vertical-align: middle;
}

.modal-container {
  width: 300px;
  margin: 0px auto;
  padding: 20px 30px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
  font-family: Helvetica, Arial, sans-serif;
}

.modal-header h3 {
  margin-top: 0;
  color: #42b983;
}

.modal-body {
  margin: 20px 0;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter {
  opacity: 0;
}

.modal-leave-active {
  opacity: 0;
}

.modal-enter .modal-container,
.modal-leave-active .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
.error {
  color: red;
}
</style>
