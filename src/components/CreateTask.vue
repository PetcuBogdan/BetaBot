<template>
  <form @submit="onSubmit" class="add-form">
    <div class="form-control">
      <label>Task name</label>
      <input
        required
        type="text"
        v-model="task.name"
        name="TaskName"
        placeholder="Task name"
      />
    </div>
    <div class="form-control">
      <label>Shop</label>
      <select
        class="btnCard"
        v-model="task.shop"
        required
        @change="getCategory"
      >
        <option value="Palace">Palace</option>
        <option value="Supreme">Supreme</option>
      </select>
    </div>
    <div class="form-control">
      <label>Category</label>
      <select
        class="btnCard"
        v-model="task.category"
        @change="getProducts"
        required
      >
        <option
          v-for="category in categories"
          :key="category"
          :value="category"
        >
          {{ category }}
        </option>
      </select>
    </div>
    <div class="form-control">
      <label>Product name</label>
      <select
        class="btnCard"
        v-model="task.productName"
        required
        @change="getSizeColor"
      >
        <option
          v-for="product in products"
          :key="product"
          :value="product.productName"
        >
          {{ product.productName }}
        </option>
      </select>
    </div>
    <div class="form-control">
      <label>Size </label>
      <select class="btnCard" v-model="task.size" required>
        <option v-for="size in sizes" :key="size" :value="size">
          {{ size }}
        </option>
      </select>
    </div>
    <div class="form-control">
      <label>Color</label>
      <select class="btnCard" v-model="task.color" required>
        <option v-for="color in colors" :key="color" :value="color">
          {{ color }}
        </option>
      </select>
    </div>
    <div class="form-control">
      <label>Profile</label>
      <select class="btnCard" v-model="task.profile" required>
        <option value="0" disabled>Select profile ...</option>
        <option
          v-for="profile in profiles"
          :key="profile.id"
          :value="profile.id"
        >
          {{ profile.fname }}
        </option>
      </select>
    </div>
    <div class="form-control">
      <label>Card</label>
      <select class="btnCard" v-model="task.card" required>
        <option value="0" disabled>Select card ...</option>
        <option :key="card.id" v-for="card in cards" :value="card.id">
          {{ card.nameCard }}
        </option>
      </select>
    </div>
    <input type="submit" value="Save Task" class="btn btnSave" />
  </form>
</template>

<script>
import axios from "axios";

var IdTaskCount = 1;

export default {
  name: "CreateTask",
  props: {
    cards: Array,
    profiles: Array,
  },
  data() {
    return {
      task: {
        id: "",
        name: "",
        shop: "",
        productName: "",
        category: "",
        size: "",
        color: "",
        profile: "",
        card: "",
        status: "",
      },
      products: [],
      supremeProducts: [],
      palaceProducts: [],
      categories: [],
      sizes: [],
      colors: [],
    };
  },
  methods: {
    getSizeColor() {
      this.sizes = [];
      this.colors = [];
      if (this.task.shop == "Palace") {
        this.colors.push("Color already selected with product");
        if (this.task.category == "HATS") {
          this.sizes = ["S/M", "L/XL"];
        } else if (this.task.category == "FOOTWEAR") {
          this.sizes = [
            "UK 6 / US 7",
            "UK 6.5 / US 7.5",
            "UK 7 / US 8",
            "UK 7.5 / US 8.5",
            "UK 8 / US 9",
            "UK 8.5 / US 9.5",
            "UK 9 / US 10",
            "UK 9.5 / US 10.5",
            "UK 10 / US 11",
            "UK 10.5 / US 11.5",
            "UK 11 / US 12",
          ];
        } else if (
          this.task.category == "ACCESSORIES" ||
          this.task.category == "HARDWARE" ||
          this.task.category == "BAGS"
        ) {
          this.sizes.push("One Size");
        } else {
          this.sizes = ["Small", "Medium", "Large", "X-Large"];
        }
      } else {
        if (
          this.task.category == "hats" ||
          this.task.category == "accessories" ||
          this.task.category == "skate"
        ) {
          this.sizes.push("One Size");
        } else {
          this.sizes = ["Small", "Medium", "Large", "XLarge"];
        }
        for (const product of this.supremeProducts) {
          if (
            product.productName.toLowerCase() ==
            this.task.productName.toLowerCase()
          ) {
            this.colors = product.colors.split(" ");
          }
        }
      }
    },
    getProducts() {
      this.task.productName = "";
      this.sizes = [];
      this.colors = [];
      let list =
        this.task.shop == "Palace" ? this.palaceProducts : this.supremeProducts;
      this.products = list.filter(
        (product) =>
          product.category.toLowerCase() === this.task.category.toLowerCase()
      );
    },
    getCategory() {
      this.task.product = [];
      this.sizes = [];
      this.colors = [];
      this.products = [];
      if (this.task.shop == "Palace") {
        this.categories = [
          "CK1 PALACE",
          "JACKETS",
          "SHIRTING",
          "TOPS",
          "BOTTOMS",
          "TRSCKSUITS",
          "HOODS",
          "SWEATSHIRTS",
          "T-SHIRTS",
          "HATS",
          "FOOTWEAR",
          "BAGS",
          "ACCESSORIES",
          "HARDWARE",
        ];
      } else {
        this.categories = [
          "jackets",
          "shirts",
          "tops/sweaters",
          "pants",
          "t-shirts",
          "hats",
          "accessories",
          "skate",
        ];
      }
    },
    onSubmit(e) {
      e.preventDefault();

      this.$emit("create-task", {
        ...this.task,
        id: IdTaskCount,
        status: "paused",
      });
      this.task.id = IdTaskCount;
      this.task.status = "paused";

      const axios = require("axios").default;

      axios
        .post("http://localhost:9000/taskAdd", this.task)
        .then(function (response) {
          console.log(response);
        })
        .catch(function (error) {
          console.log(error);
        });

      IdTaskCount += 1;

      this.task = {
        id: "",
        name: "",
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
    async initTaskId() {
      let response = await axios.get("http://localhost:9000/taskId");
      IdTaskCount = response?.data;
      IdTaskCount++;
    },
    async initPalaceProducts() {
      let response = await axios.get("http://localhost:9000/getPalaceProducts");
      this.palaceProducts = [...response?.data];
      //console.log(this.palaceProducts);
    },
    async initSupremeProducts() {
      let response = await axios.get(
        "http://localhost:9000/getSupremeProducts"
      );
      this.supremeProducts = [...response?.data];
      //console.log(this.supremeProducts);
    },
  },
  async created() {
    this.initTaskId();
    this.initSupremeProducts();
    this.initPalaceProducts();
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
.btnCard {
  display: inline-block;
  background: white;
  color: black;
  border: none;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
  text-decoration: none;
  font-size: 15px;
  font-family: inherit;
  display: block;
  width: 310px;
  height: 40px;
  margin-bottom: 15px;
  margin-top: 5px;
}
</style>
