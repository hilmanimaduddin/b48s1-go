console.log("aaaaaaaaaaa");

const promise = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest();

  xhr.open("GET", "https://api.npoint.io/36c88be3ce02946fb58a", true);
  xhr.onload = function () {
    if (xhr.status === 200) {
      resolve(JSON.parse(xhr.responseText));
    } else if (xhr.status >= 400) {
      reject("Error loading data");
    }
  };

  xhr.onerror = function () {
    reject("Network error");
  };

  xhr.send();
});

// promise
//   .then((value) => {
//     console.log(value);
//   })
//   .catch((uuu) => {
//     console.log(uuu);
//   });

let testimonialData = [];

async function getData(rating) {
  try {
    const response = await promise;
    console.log(response);
    testimonialData = response;
    allTestimonial();
  } catch (err) {
    console.log(er);
  }
}

getData();

function allTestimonial() {
  let testimonialHTML = "";

  testimonialData.forEach((card, index) => {
    testimonialHTML += `<div class="card">
                          <img
                          src="${card.image}"
                          alt="foto"/>
                          <p class="quote">"${card.quote}"</p>
                          <p class="author">~ ${card.user}</p>
                          <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
                      </div>`;
  });
  document.getElementById("testimonials").innerHTML = testimonialHTML;
}

function filterTestimonial(rating) {
  let filterTestimonialHtml = "";

  const filteredData = testimonialData.filter((card) => {
    return card.rating === rating;
  });

  filteredData.forEach((card) => {
    filterTestimonialHtml += `<div class="card">
                                <img
                                src="${card.image}"
                                alt="foto"/>
                                <p class="quote">"${card.quote}"</p>
                                <p class="author">~ ${card.user}</p>
                                <p class="author">${card.rating} <i class="fa-solid fa-star"></i></p>
                            </div>`;
  });
  if (filterTestimonialHtml == "") {
    filterTestimonialHtml = `<h1>Nothing</h1>`;
  }

  document.getElementById("testimonials").innerHTML = filterTestimonialHtml;
}
