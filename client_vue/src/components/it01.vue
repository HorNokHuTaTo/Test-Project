<template>
  <div>
    <header class="header">IT 01-1</header>
    <section class="container">
      <button class="add-btn" @click="showAdd = true">ADD</button>
      <table class="data-table">
        <thead>
          <tr>
            <th>Id</th>
            <th>ชื่อ-สกุล</th>
            <th>ที่อยู่</th>
            <th>วันเกิด</th>
            <th>อายุ</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(person, idx) in persons" :key="person.id">
             <td>{{ idx + 1 }}</td>
            <td>{{ person.full_name }}</td>
            <td>{{ person.address }}</td>
            <td>{{ format(person.birth_date) }}</td>
            <td>{{ person.age }}</td>
            <td>
              <button class="view-btn" @click="view(person)">View</button>
            </td>
          </tr>
        </tbody>
      </table>
    </section>

    <div v-if="showAdd" class="modal-bg">
      <div class="modal-box">
        <header class="header">IT 01-2</header>
        <form @submit.prevent="addPerson">
          <div class="form-row">
            <label>ชื่อ - สกุล</label>
            <input v-model="form.full_name" required />
          </div>
          <div class="form-row">
            <label>วันเกิด</label>
            <input type="date" v-model="form.birth_date" @change="updateAge" required />
          </div>
          <div class="form-row">
            <label>อายุ</label>
            <span>{{ form.age ? form.age : 'xx' }} ปี</span>
          </div>
          <div class="form-row">
            <label>ที่อยู่</label>
            <textarea v-model="form.address" required></textarea>
          </div>
          <div class="form-actions">
            <button type="submit" class="save-btn">บันทึก</button>
            <button type="button" class="cancel-btn" @click="closeAdd">ยกเลิก</button>
          </div>
        </form>
      </div>
    </div>

    <div v-if="showView" class="modal-bg">
      <div class="modal-box">
        <header class="header">IT 01-3</header>
        <div class="form-row">
          <label>ชื่อ - สกุล</label>
          <input :value="selected.full_name" readonly />
        </div>
        <div class="form-row">
          <label>วันเกิด</label>
          <input :value="format(selected.birth_date)" readonly />
        </div>
        <div class="form-row">
          <label>อายุ</label>
          <span>{{ selected.age ? selected.age : 'xx' }} ปี</span>
        </div>
        <div class="form-row">
          <label>ที่อยู่</label>
          <textarea :value="selected.address" readonly></textarea>
        </div>
        <div class="form-actions">
          <button class="close-btn" @click="showView = false">ปิด</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'IT01',
  data() {
    return {
      persons: [],
      showAdd: false,
      showView: false,
      form: { full_name: '', birth_date: '', age: '', address: '' },
      selected: {},
      apiBase: 'https://test-project-0w71.onrender.com/api/it01/people', 
    }
  },
  mounted() {
    this.fetchPersons()
  },
  methods: {
    format(date) {
      if (!date) return ''
      const d = new Date(date)
      return `${String(d.getDate()).padStart(2, '0')}/${String(d.getMonth() + 1).padStart(2, '0')}/${d.getFullYear()}`
    },
    updateAge() {
  if (this.form.birth_date) {
    const birthYear = Number(this.form.birth_date.split('-')[0])
    this.form.age = new Date().getFullYear() - birthYear
  } else {
    this.form.age = ''
  }
},
    async fetchPersons() {
      try {
        const res = await fetch(this.apiBase)
        this.persons = await res.json()
      } catch (e) {
        alert('โหลดข้อมูลไม่สำเร็จ')
      }
    },
    async addPerson() {
      try {
        const res = await fetch(this.apiBase, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            full_name: this.form.full_name,
            birth_date: this.form.birth_date,
            address: this.form.address,
          }),
        })
        if (!res.ok) throw new Error('เพิ่มข้อมูลไม่สำเร็จ')
        const person = await res.json()
        this.persons.push({
          id: person.id,
          full_name: person.full_name,
          birth_date: person.birth_date,
          age: person.age,
          address: person.address,
        })
        this.closeAdd()
      } catch (e) {
        alert(e.message)
      }
    },
    closeAdd() {
      this.showAdd = false
      this.form = { full_name: '', birth_date: '', age: '', address: '' }
    },
    async view(person) {
      try {
        const res = await fetch(`${this.apiBase}/${person.id}`)
        if (!res.ok) throw new Error('ไม่พบข้อมูล')
        const p = await res.json()
        this.selected = {
          id: p.id,
          full_name: p.full_name,
          birth_date: p.birth_date,
          age: p.age,
          address: p.address,
        }
        this.showView = true
      } catch (e) {
        alert(e.message)
      }
    },
  },
}
</script>
