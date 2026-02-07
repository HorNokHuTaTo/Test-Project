<template>
  <div class="form-container">
    <div class="form-header">IT 04-1</div>
    <div v-if="successMessage" class="success-popup">
      {{ successMessage }}
        </div>
    <form @submit.prevent="submitForm">
      <!-- First Row -->
      <div class="form-row">
        <div class="form-group">
          <label>First Name</label>
          <input 
            v-model="formData.firstName" 
            type="text"
            @blur="validateField('firstName')"
          />
          <span v-if="errors.firstName" class="error-text">{{ errors.firstName }}</span>
        </div>
        <div class="form-group">
          <label>Last Name</label>
          <input 
            v-model="formData.lastName" 
            type="text"
            @blur="validateField('lastName')"
          />
          <span v-if="errors.lastName" class="error-text">{{ errors.lastName }}</span>
        </div>
      </div>

      <!-- Second Row -->
      <div class="form-row">
        <div class="form-group">
          <label>Email</label>
          <input 
            v-model="formData.email" 
            type="email"
            @blur="validateField('email')"
          />
          <span v-if="errors.email" class="error-text">{{ errors.email }}</span>
        </div>
        <div class="form-group">
          <label>Phone</label>
          <input 
            v-model="formData.phone" 
            type="text"
            @blur="validateField('phone')"
          />
          <span v-if="errors.phone" class="error-text">{{ errors.phone }}</span>
        </div>
      </div>

      <!-- Third Row -->
      <div class="form-row">
        <div class="form-group">
          <label>Profile</label>
          <div class="input-with-button">
            <input 
              ref="profileName"
              v-model="formData.profileName" 
              type="text"
              placeholder="No file chosen"
              readonly
            />
            <input 
              ref="profileInput"
              @change="handleProfileChange"
              type="file" 
              accept="any/*"
              style="display: none"
            />
            <button type="button" class="browse-btn" @click="$refs.profileInput.click()">Browse</button>
          </div>
          <span v-if="errors.profile" class="error-text">{{ errors.profile }}</span>
        </div>
        <div class="form-group">
          <label>Birth Day</label>
          <div class="input-with-button">
            <input v-model="formData.birthDay" type="date" />
            <!-- <button type="button" class="date-btn">📅</button> -->
          </div>
          <span v-if="errors.birthDay" class="error-text">{{ errors.birthDay }}</span>
        </div>
      </div>

      <!-- Fourth Row -->
      <div class="form-row">
        <div class="form-group full-width">
          <label>Occupation</label>
          <select 
            v-model="formData.occupation"
            @blur="validateField('occupation')"
          >
            <option value="">Select Occupation</option>
            <option v-for="occ in occupations" :key="occ.id" :value="occ.id">
              {{ occ.label }}
            </option>
          </select>
          <span v-if="errors.occupation" class="error-text">{{ errors.occupation }}</span>
        </div>
      </div>

      <!-- Sex Radio -->
      <div class="form-row">
        <div class="form-group full-width">
          <label>Sex</label>
          <div class="radio-group">
            <label>
              <input v-model="formData.sex" type="radio" value="male" />
              Male
            </label>
            <label>
              <input v-model="formData.sex" type="radio" value="female" />
              Female
            </label>
          </div>
        </div>
      </div>

      <!-- Buttons -->
      <div class="form-buttons">
        <button type="submit" class="save-btn">Save</button>
        <button type="button" @click="clearForm" class="clear-btn">Clear</button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  name: 'IT04',
  data() {
    return {
      formData: {
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        profileFile: null,
        profileName: '',
        birthDay: '',
        occupation: '',
        sex: ''
      },
      occupations: [],
      errors: {},
      successMessage: ''
    }
  },
  mounted() {
    this.fetchOccupations();
  },
  methods: {
    async fetchOccupations() {
      try {
        const res = await fetch('http://localhost:3000/api/it04/occupations');
        if (!res.ok) throw new Error('Failed to fetch occupations');
        this.occupations = await res.json();
      } catch (err) {
        console.error(err);
      }
    },
    validateField(fieldName) {
      switch(fieldName) {
        case 'firstName':
          if (!this.formData.firstName) this.errors.firstName = 'First Name is required'; else delete this.errors.firstName;
          break;
        case 'lastName':
          if (!this.formData.lastName) this.errors.lastName = 'Last Name is required'; else delete this.errors.lastName;
          break;
        case 'email':
          if (!this.formData.email) this.errors.email = 'Email is required';
          else if (!this.isValidEmail(this.formData.email)) this.errors.email = 'Please provide a valid Email';
          else delete this.errors.email;
          break;
        case 'phone':
          if (!this.formData.phone) this.errors.phone = 'Phone is required';
          else if (!this.isValidPhone(this.formData.phone)) this.errors.phone = 'Please provide a valid Phone';
          else delete this.errors.phone;
          break;
        case 'profile':
          if (!this.formData.profileFile) this.errors.profile = 'Please selected any profile'; else delete this.errors.profile;
          break;
        case 'birthDay':
          if (!this.formData.birthDay) this.errors.birthDay = 'Please provide a valid Birth Day'; else delete this.errors.birthDay;
          break;
        case 'occupation':
          if (!this.formData.occupation) this.errors.occupation = 'Please selected any Occupation'; else delete this.errors.occupation;
          break;
        case 'sex':
          if (!this.formData.sex) this.errors.sex = 'Sex is required'; else delete this.errors.sex;
          break;
      }
    },
    isValidEmail(email) {
      const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      return re.test(email);
    },
    isValidPhone(phone) {
      const re = /^0[0-9]{9}$/; // example: Thai format 10 digits starting with 0
      return re.test(phone);
    },
    handleProfileChange(e) {
      const file = e.target.files[0];
      if (file) {
        this.formData.profileFile = file;
        this.formData.profileName = file.name;
        this.validateField('profile');
      }
    },
    formatDateToDDMMYYYY(isoDate) {
      if (!isoDate) return '';
      const d = new Date(isoDate);
      const dd = String(d.getDate()).padStart(2, '0');
      const mm = String(d.getMonth() + 1).padStart(2, '0');
      const yyyy = d.getFullYear();
      return `${dd}/${mm}/${yyyy}`;
    },
    async submitForm() {
      // validate all
      ['firstName','lastName','email','phone','profile','birthDay','occupation','sex'].forEach(f => this.validateField(f));
      if (Object.keys(this.errors).length) return;

      try {
        const fd = new FormData();
        fd.append('first_name', this.formData.firstName);
        fd.append('last_name', this.formData.lastName);
        fd.append('email', this.formData.email);
        fd.append('phone', this.formData.phone);
        // backend expects dd/mm/yyyy
        fd.append('birthday', this.formatDateToDDMMYYYY(this.formData.birthDay));
        fd.append('occupation', this.formData.occupation);
        fd.append('sex', this.formData.sex);
        if (this.formData.profileFile) fd.append('profile', this.formData.profileFile);

        const res = await fetch('http://localhost:3000/api/it04/insert-profile', {
          method: 'POST',
          body: fd
        });

        const data = await res.json();
        if (!res.ok) {
          this.successMessage = data.message || data.error || 'Save failed';
          setTimeout(() => { this.successMessage = ''; }, 3000);
          return;
        }
 
        
        // ensure message includes id
        this.successMessage = `${data.message} id : ${data.id}`;
        setTimeout(()=> this.successMessage = '', 3000);
        // clear form after success
        this.clearForm();
      } catch (err) {
        console.error(err);
        this.successMessage = 'Save failed';
        setTimeout(() => { this.successMessage = ''; }, 3000);
      }
    },
    clearForm() {
      this.formData = {
        firstName: '',
        lastName: '',
        email: '',
        phone: '',
        profileFile: null,
        profileName: '',
        birthDay: '',
        occupation: '',
        sex: ''
      };
      this.errors = {};
      this.message = '';
      if (this.$refs.profileInput) this.$refs.profileInput.value = '';
      if (this.$refs.profileName) this.$refs.profileName.value = '';
    }
  }
}
</script>

<style scoped>
.form-container {
  max-width: 800px;
  margin: 20px auto;
  border: 1px solid #ddd;
  border-radius: 4px;
  overflow: visible;
}

.form-header {
  background-color: #27ae60;
  color: white;
  padding: 12px;
  font-weight: bold;
  font-size: 16px;
    text-align: center;
}

form {
  padding: 20px;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  margin-bottom: 20px;
  position: relative;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-group.full-width {
  grid-column: 1 / -1;
}

label {
  margin-bottom: 6px;
  font-weight: 500;
  font-size: 13px;
}

input[type="text"],
input[type="email"],
input[type="date"],
select {
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 13px;
  font-family: Arial, sans-serif;
}

input[type="text"]:focus,
input[type="email"]:focus,
input[type="date"]:focus,
select:focus {
  outline: none;
  border-color: #4CAF50;
  box-shadow: 0 0 4px rgba(76, 175, 80, 0.3);
}

.input-with-button {
  display: flex;
  gap: 8px;
}

.input-with-button input {
  flex: 1;
}

.browse-btn,
.date-btn {
  padding: 8px 12px;
  background-color: #999;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 13px;
}

.browse-btn:hover,
.date-btn:hover {
  background-color: #777;
}

.radio-group {
  display: flex;
  gap: 20px;
  margin-top: 6px;
}

.radio-group label {
  display: flex;
  align-items: center;
  margin-bottom: 0;
  gap: 6px;
}

.radio-group input[type="radio"] {
  margin: 0;
}

.error-text {
  color: #d32f2f;
  font-size: 12px;
  margin-top: 4px;
}

.success-popup {
  position: fixed;
  top: 20px;
  right: 20px;
  background-color: #27ae60;
  color: white;
  padding: 16px 24px;
  border-radius: 4px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  font-size: 14px;
  font-weight: 500;
  z-index: 9999;
  animation: slideIn 0.3s ease-out;
}

@keyframes slideIn {
  from {
    transform: translateX(400px);
    opacity: 0;
  }
  to {
    transform: translateX(0);
    opacity: 1;
  }
}

.form-buttons {
  display: flex;
  gap: 10px;
  justify-content: center;
  margin-top: 30px;
}

.save-btn {
  background-color: #4CAF50;
  color: white;
  padding: 10px 30px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
}

.save-btn:hover {
  background-color: #45a049;
}

.clear-btn {
  background-color: #999;
  color: white;
  padding: 10px 30px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
}

.clear-btn:hover {
  background-color: #777;
}
</style>