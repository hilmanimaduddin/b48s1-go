// class Testimonial {
//   #quote = "";
//   #image = "";

//   constructor(quote, image) {
//     this.#quote = quote;
//     this.#image = image;
//   }

//   get quote() {
//     return this.#quote;
//   }

//   get image() {
//     return this.#image;
//   }

//   get user() {
//     return new Error(`there is must be user to make testimonials`);
//   }

//   get testimonialHTML() {
//     return `<div class="card">
//                 <img
//                 src="${this.image}"
//                 alt="foto"
//                 />
//                 <p class="quote">"${this.quote}"</p>
//                 <p class="author">~ ${this.user}</p>
//             </div>`;
//   }
// }

// class UserTestimonial extends Testimonial {
//   #user = "";

//   constructor(user, quote, image) {
//     super(quote, image);
//     this.#user = user;
//   }

//   get user() {
//     return "User : " + this.#user;
//   }
// }

// class CompanyTestimonial extends Testimonial {
//   #company = "";

//   constructor(company, quote, image) {
//     super(quote, image);
//     this.#company = company;
//   }

//   get user() {
//     return "Company : " + this.#company;
//   }
// }

// const testimonial1 = new UserTestimonial(
//   "Cak Jon",
//   "Pak leknya Joko",
//   "https://images.unsplash.com/photo-1685242069669-dfcaf86c69eb?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=327&q=80"
// );

// const testimonial2 = new CompanyTestimonial(
//   "Suka Maju",
//   "Harap agar tetap maju",
//   "https://images.unsplash.com/photo-1516959543587-4cc499d9514b?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTN8fGluZG9uZXNpYSUyMHBlb3BsZXxlbnwwfHwwfHx8Mg%3D%3D&auto=format&fit=crop&w=600&q=60"
// );

// const testimonial3 = new UserTestimonial(
//   "Anak Bolang",
//   "Dulu main sampe lupa waktu, sekarang maen kalo ada waktu",
//   "https://images.unsplash.com/photo-1502781252888-9143ba7f074e?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8aW5kb25lc2lhJTIwcGVvcGxlfGVufDB8fDB8fHwy&auto=format&fit=crop&w=600&q=60"
// );

// let testimonialData = [testimonial1, testimonial2, testimonial3];

// let testimonialHTML = "";

// for (let i = 0; i < testimonialData.length; i++) {
//   testimonialHTML += testimonialData[i].testimonialHTML;
// }

// document.getElementById("testimonials").innerHTML = testimonialHTML;

const testimonialData = [
  {
    user: "Cak Jon",
    quote: "Pak leknya Joko",
    image:
      "https://images.unsplash.com/photo-1685242069669-dfcaf86c69eb?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=327&q=80",
    rating: 3,
  },
  {
    user: "Pak Joko",
    quote: "Mau cukur gak?",
    image:
      "https://images.unsplash.com/photo-1514178320695-089eff7b8aa8?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MjB8fGluZG9uZXNpYSUyMHBlb3BsZXxlbnwwfHwwfHx8Mg%3D%3D&auto=format&fit=crop&w=600&q=60",
    rating: 5,
  },
  {
    user: "Mukidi",
    quote: "Main yuk, ke sawah!!",
    image:
      "https://images.unsplash.com/photo-1502781252888-9143ba7f074e?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8aW5kb25lc2lhJTIwcGVvcGxlfGVufDB8fDB8fHwy&auto=format&fit=crop&w=600&q=60",
    rating: 2,
  },
  {
    user: "Bocilll",
    quote: "Rodanya berputar..",
    image:
      "https://images.unsplash.com/photo-1516959543587-4cc499d9514b?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTN8fGluZG9uZXNpYSUyMHBlb3BsZXxlbnwwfHwwfHx8Mg%3D%3D&auto=format&fit=crop&w=600&q=60",
    rating: 3,
  },
  {
    user: "Empat Sekawan",
    quote: "Bolanya kelempar ke atas",
    image:
      "https://images.unsplash.com/photo-1502086223501-7ea6ecd79368?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTd8fGluZG9uZXNpYSUyMHBlb3BsZXxlbnwwfHwwfHx8Mg%3D%3D&auto=format&fit=crop&w=600&q=60",
    rating: 5,
  },
  {
    user: "Sri",
    quote: "Persiapan Buat Lomba..",
    image:
      "https://images.unsplash.com/photo-1599717462650-02831c496816?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MTl8fGluZG9uZXNpYSUyMHBlb3BsZXxlbnwwfHwwfHx8Mg%3D%3D&auto=format&fit=crop&w=600&q=60",
    rating: 1,
  },
  {
    user: "Anak Senja",
    quote: "Naik sepeda sore emang enak",
    image:
      "https://images.unsplash.com/photo-1442544213729-6a15f1611937?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8aW5kb25lc2lhfGVufDB8fDB8fHwy&auto=format&fit=crop&w=600&q=60",
    rating: 3,
  },
  {
    user: "Kang Becak",
    quote: "Butuh Tumpangan??",
    image:
      "https://images.unsplash.com/photo-1527965408463-82ae0731825c?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8OXx8aW5kb25lc2lhfGVufDB8fDB8fHwy&auto=format&fit=crop&w=600&q=60",
    rating: 2,
  },
];

console.log(testimonialData);

function allTestimonial() {
  let testimonialHTML = "";

  testimonialData.forEach((value, index) => {
    testimonialHTML += `<div class="card">
                          <img
                          src="${value.image}"
                          alt="foto"/>
                          <p class="quote">"${value.quote}"</p>
                          <p class="author">~ ${value.user}</p>
                          <p class="author">${value.rating} <i class="fa-solid fa-star"></i></p>
                      </div>`;
  });

  document.getElementById("testimonials").innerHTML = testimonialHTML;
}

allTestimonial();

function filterTestimonial(rating) {
  let filteredTestimonialHTML = "";

  const filteredData = testimonialData.filter((value) => {
    return value.rating === rating;
  });

  filteredData.forEach((value) => {
    filteredTestimonialHTML += `<div class="card">
                                  <img
                                  src="${value.image}"
                                  alt="foto"/>
                                  <p class="quote">"${value.quote}"</p>
                                  <p class="author">~ ${value.user}</p>
                                  <p class="author">${value.rating} <i class="fa-solid fa-star"></i></p>
                              </div>`;
  });

  if (filteredTestimonialHTML == "") {
    filteredTestimonialHTML = `<h1>Nothing</h1>`;
  }

  document.getElementById("testimonials").innerHTML = filteredTestimonialHTML;
}
