<template>
  <div>
    <div v-if="page === 'login'" class="box">
      
      <header class="header">IT 02-1</header>
      <div class="form-row">
        <label>Username</label>
        <input v-model="loginForm.username" />
      </div>
      <div class="form-row">
        <label>Password</label>
        <input v-model="loginForm.password" type="password" />
      </div>
      <div class="form-actions">
        <button @click="doLogin" class="save-btn">ลงชื่อเข้าใช้งาน</button>
        <div v-if="loginError" style="color:red; text-align:center;">{{ loginError }}</div>
      </div>
      <div class="form-actions">
        <a href="#" @click.prevent="page = 'register'">สมัครสมาชิก</a>
      </div>
    </div>

    <!-- IT 02-2: Register Form -->
    <div v-if="page === 'register'" class="box">
      <header class="header">IT 02-2</header>
      <div class="form-row">
        <label>Username</label>
        <input v-model="registerForm.username" />
      </div>
      <div class="form-row">
        <label>Password</label>
        <input v-model="registerForm.password" type="password" />
      </div>
      <div class="form-row">
        <label>Confirm Password</label>
        <input v-model="registerForm.confirm_password" type="password" />
      </div>
      <div class="form-actions">
        <button @click="doRegister" class="save-btn">สมัครสมาชิก</button>
      </div>
      <div class="form-actions">
    <button @click="goToLogin" class="save-btn" style="margin-top:10px;">กลับหน้าแรก</button>
  </div>
      <div v-if="registerError" style="color:red; text-align:center;">{{ registerError }}</div>
    </div>

    <!-- IT 02-3: Welcome -->
    <div v-if="page === 'welcome'" class="box">
      <header class="header">IT 02-3</header>
      <div style="text-align:center; margin-top:60px;">
        Welcome User : {{ currentUser }}
      </div>
      <div class="form-actions" style="margin-top:30px;">
    <button @click="goToLogin" class="save-btn">OK</button>
  </div>
    </div>
  </div>
</template>

<script>
async function registerUser(username, password) {
  const res = await fetch('http://localhost:3000/api/it02/register', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, password, confirm_password: password })
  });
  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || 'Register failed');
  }
  return await res.json();
}

async function loginUser(username, password) {
  const res = await fetch('http://localhost:3000/api/it02/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ username, password })
  });
  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || 'Login failed');
  }
  return await res.json();
}

export default {
  name: "IT02",
  data() {
    return {
      page: "login", 
      loginForm: { username: "", password: "" },
      registerForm: { username: "", password: "", confirm_password: "" },
      registerError: "",
      currentUser: "",
      token: '',
      loginError: '',
    };
  },
  methods: {
    async doRegister() {
      this.registerError = '';
      if (!this.registerForm.username || !this.registerForm.password || !this.registerForm.confirm_password) {
        this.registerError = 'กรุณากรอกข้อมูลให้ครบ';
        return;
      }
      if (this.registerForm.password !== this.registerForm.confirm_password) {
        this.registerError = 'Password ไม่ตรงกัน';
        return;
      }
      try {
        await registerUser(this.registerForm.username, this.registerForm.password);
        alert('สมัครสมาชิกสำเร็จ');
        this.page = "login";
        this.registerForm = { username: '', password: '', confirm_password: '' };
      } catch (e) {
        this.registerError = e.message;
      }
    },
    async doLogin() {
      this.loginError = '';
      if (!this.loginForm.username || !this.loginForm.password) {
        this.loginError = 'กรุณากรอก User และ Password';
        return;
      }
      try {
        const result = await loginUser(this.loginForm.username, this.loginForm.password);
        this.token = result.token;
        this.currentUser = result.username;
        this.page = "welcome";
      } catch (e) {
        this.loginError = e.message;
      }
    },
    goToLogin() {
    this.page = "login";
    this.currentUser = "";
    this.loginForm = { username: '', password: '' };
    this.token = '';
    this.loginError = '';
    }
  }
}
</script>

<style scoped>
.box {
  max-width: 400px;
  margin: 30px auto;
  border: 2px solid #222;
  background: #fff;
  padding-bottom: 30px;
  border-radius: 8px;
}
.header {
  background: #27ae60;
  color: #fff;
  text-align: center;
  padding: 10px;
  font-weight: bold;
  border-radius: 6px 6px 0 0;
}
.form-row {
  margin: 30px 30px 0 30px;
  display: flex;
  align-items: center;
}
.form-row label {
  width: 100px;
}
.form-row input {
  margin-left: 10px;
  width: 200px;
  padding: 4px;
}
.form-actions {
  margin: 20px 30px 0 30px;
  text-align: center;
}
.save-btn {
  background: #ecf0f1;
  color: #222;
  border: 1px solid #bbb;
  padding: 8px 16px;
  border-radius: 4px;
  font-weight: bold;
  cursor: pointer;
  transition: background 0.2s;
}
.save-btn:hover {
  background: #bdc3c7;
}
a {
  color: #2980b9;
  text-decoration: underline;
  cursor: pointer;
}
</style>