<template>
  <form @submit="onSubmit" class="add-form">
    <div class="form-control">
      <label>Task name</label>
      <input
        type="text"
        v-model="task"
        name="TaskName"
        placeholder="Task name"
      />
    </div>
    <div class="form-control">
      <label>Shop</label>
      <input type="text" v-model="shop" name="Shop" placeholder="Shop" />
    </div>
    <div class="form-control">
      <label>Product name</label>
      <input
        type="text"
        v-model="product"
        name="ProductName"
        placeholder="Product name"
      />
    </div>
    <div class="form-control">
      <label>Category</label>
      <input
        type="text"
        v-model="category"
        name="category"
        placeholder="Category"
      />
    </div>
    <div class="form-control">
      <label>Size </label>
      <input
        type="text"
        v-model="size"
        name="size"
        placeholder="Example: S/M/L/XL"
      />
    </div>
    <div class="form-control">
      <label>Color</label>
      <input type="text" v-model="color" name="color" placeholder="Color" />
    </div>
    <div class="form-control">
      <label>Profile</label>
      <select class="drop" v-model="profile">
        <option
          value="0"
          label="Select profile... "
          selected
          class="selectDefat"
        >
          Select profile ...
        </option>
        <option value="Profile1">Profile1</option>
        <option value="Profile2">Profile2</option>
        <option value="Profile3">Profile3</option>
      </select>
    </div>
    <div class="form-control">
      <label>Card</label>
      <select class="drop" v-model="card">
        <option value="0" label="Select Card... " selected class="selectDefat">
          Select card ...
        </option>
        <option value="Card1">Card1</option>
        <option value="Card2">Card2</option>
        <option value="Card3">Card3</option>
      </select>
    </div>
    <input type="submit" value="Save Task" class="btn btnSave" />
  </form>
</template>

<script>
export default {
  name: "CreateTask",
  data() {
    return {
      id: "",
      taskName: "",
      shop: "",
      productName: "",
      category: "",
      size: "",
      color: "",
      profile: "",
      card: "",
      status: "",
    };
  },
  methods: {
    onSubmit(e) {
      e.preventDefault();
      if (!this.task) {
        alert("Please add name for task");
        return;
      }
      if (!this.shop) {
        alert("Please add shop name");
        return;
      }
      if (!this.product) {
        alert("Please add product name");
        return;
      }
      if (!this.category) {
        alert("Please add product category");
        return;
      }
      if (!this.profile) {
        alert("Please select profile");
        return;
      }
      if (!this.card) {
        alert("Please add card");
        return;
      }

      const newTask = {
        id: Math.floor(Math.random() * 100000),
        task: this.task,
        shop: this.shop,
        product: this.product,
        category: this.category,
        size: this.size,
        color: this.color,
        profile: this.profile,
        card: this.card,
        status: "paused",
        reminder: true,
      };

      this.$emit("create-task", newTask);

      this.shop = "";
      this.task = "";
      this.product = "";
      this.category = "";
      this.size = "";
      this.color = "";
      this.profile = "";
      this.card = "";
    },
  },
};
</script>

<style scoped>
.add-form {
  margin-bottom: 40px;
}
.form-control {
  float: left;
  margin: 20px 100px;
  margin-top: 0px;
}
.form-control label {
  display: block;
  color: cadetblue;
}
.form-control input {
  width: 300px;
  height: 35px;
  margin: 5px;
  padding: 3px 7px;
  font-size: 17px;
}
.form-control-check label {
  flex: 1;
  color: cadetblue;
}
.form-control-check input {
  flex: 2;
  height: 20px;
}
.btnSave {
  width: 300px;
  height: 40px;
  margin: 5px;
  padding: 3px 7px;
  margin-left: 100px;
}
.drop {
  width: 300px;
  height: 34px;
  display: inline-block;
  background: white;
  color: black;
  border: none;
  padding: 3px 7px;
  margin: 5px;
  margin-top: 4px;
  border-radius: 1px;
  cursor: pointer;
  font-size: 17px;
  font-family: inherit;
}
</style>
