<template>
  <div class="page">
    <div class="header">IT 06-1</div>

    <div class="form-row">
      <label>รหัสสินค้า</label>
      <input v-model="skuInput" placeholder="XXXX-XXXX-XXXX-XXXX" />
      <button class="btn-add" @click="addSku">ADD</button>
    </div>

    <table class="sku-table">
      <thead>
        <tr>
          <th>Id</th>
          <th>รหัสสินค้า (16 หลัก)</th>
          <th>บาร์โค้ดสินค้า</th>
          <th>Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(sku, index) in skus" :key="sku.id">
          <td>{{ index + 1 }}</td>
          <td>{{ sku.sku_code }}</td>
          <td>
            <img
              :src="`data:image/png;base64,${sku.barcode}`"
              class="barcode"
            />
          </td>
          <td>
            <button class="btn-del" @click="openDelete(index)">ลบ</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-if="showModal" class="modal-backdrop">
      <div class="modal">
        <p>
          ต้องการลบข้อมูล รหัสสินค้า
          <strong>{{ deleteTarget?.sku_code }}</strong> หรือไม่ ?
        </p>
        <div class="modal-actions">
          <button class="btn-ok" @click="confirmDelete">ตกลง</button>
          <button class="btn-cancel" @click="showModal = false">ยกเลิก</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue"

const skuInput = ref("")
const skus = ref([])
const showModal = ref(false)
const deleteIndex = ref(null)
const deleteTarget = ref(null)
const API = "http://localhost:3000/api/it06" 

onMounted(async () => {
  await loadSkus()
})

async function loadSkus() {
  const res = await fetch(`${API}/skus`)
  skus.value = await res.json()
}


async function addSku() {
  if (!skuInput.value) return

  const res = await fetch(`${API}/insert-sku`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ sku_code: skuInput.value }),
  })

  if (!res.ok) {
    if ( res.status === 400 ) {
      const data = await res.json()
      alert(data.error)
      return
    }
    
    alert("Insert failed")
    return
  }

  skuInput.value = ""
  await loadSkus()
}

function openDelete(index) {
  deleteIndex.value = index
  deleteTarget.value = skus.value[index]
  showModal.value = true
}

async function confirmDelete() {
  const id = deleteTarget.value.id

  const res = await fetch(`${API}/delete-sku?id=${encodeURIComponent(id)}`, {
    method: "DELETE",
  })

  if (!res.ok) {
    alert("Delete failed")
    return
  }

  showModal.value = false
  await loadSkus()
}
</script>

<style scoped>
.page {
  width: 900px;
  margin: 20px auto;
  border: 1px solid #333;
}

.header {
  background: #00b050;
  color: #000;
  text-align: center;
  padding: 8px;
  font-weight: bold;
}

.form-row {
  padding: 10px;
  display: flex;
  gap: 10px;
  align-items: center;
}

.form-row input {
  flex: 1;
  padding: 6px;
}

.btn-add {
  background: #2e6fd3;
  color: #fff;
  border: none;
  padding: 6px 14px;
  cursor: pointer   ;
}

.sku-table {
  width: 100%;
  border-collapse: collapse;
}

.sku-table th {
  background: #00b0f0;
  border: 1px solid #333;
  padding: 6px;
}

.sku-table td {
  border: 1px solid #333;
  padding: 6px;
  text-align: center;
}

.barcode {
  height: 50px;
}

.btn-del {
  background: red;
  color: white;
  border: none;
  padding: 4px 10px;
  border-radius: 4px;
  cursor: pointer;
}

.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(0,0,0,0.3);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal {
  background: #fff;
  padding: 20px 30px;
  border-radius: 6px;
  text-align: center;
}

.modal-actions {
  margin-top: 15px;
  display: flex;
  gap: 10px;
  justify-content: center;
}

.btn-ok {
  background: #2e6fd3;
  color: white;
  border: none;
  padding: 6px 16px;
  cursor: pointer;
}

.btn-cancel {
  background: red;
  color: white;
  border: none;
  padding: 6px 16px;
  cursor: pointer;
}
</style>
