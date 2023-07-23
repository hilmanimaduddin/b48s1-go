let dataBlog = [];

// const cobaAja = (coba) => {};

function addBlog(event) {
  event.preventDefault();

  let project = document.getElementById("input-project").value;
  let startDate = new Date(document.getElementById("input-startdate").value);
  let endDate = new Date(document.getElementById("input-enddate").value);
  let description = document.getElementById("input-description").value;
  let image = document.getElementById("input-image").files;
  let time = new Date();

  image = URL.createObjectURL(image[0]);

  // checkbox
  let checkbox1 = document.getElementById("checkbox1").checked;
  let checkbox2 = document.getElementById("checkbox2").checked;
  let checkbox3 = document.getElementById("checkbox3").checked;
  let checkbox4 = document.getElementById("checkbox4").checked;

  if (checkbox1 === true) {
    checkbox1 = ` <i class="fa-brands fa-node-js fa-xl me-2" style="color: #000000"></i
    >`;
  } else {
    checkbox1 = ``;
  }

  if (checkbox2 === true) {
    checkbox2 = `<i class="fa-brands fa-react fa-xl me-2" style="color: #000000"></i
    >`;
  } else {
    checkbox2 = ``;
  }

  if (checkbox3 === true) {
    checkbox3 = `<i class="fa-brands fa-jsfiddle fa-xl me-2" style="color: #000000"></i
    >`;
  } else {
    checkbox3 = ``;
  }

  if (checkbox4 === true) {
    checkbox4 = `<i class="fa-brands fa-java fa-xl me-2" style="color: #000000"></i>`;
  } else {
    checkbox4 = ``;
  }

  let date1 = endDate.getDate();
  let month1 = endDate.getMonth();

  let date2 = startDate.getDate();
  let month2 = startDate.getMonth();

  let date = Math.floor(date1 - date2);
  let months = Math.floor(month1 - month2);

  if (date == 0) {
    date = "";
  } else if (date == 1) {
    date = date + " day";
  } else {
    date = date + " days";
  }

  if (months == 0) {
    months = "";
  } else if (months == 1) {
    months = months + " month";
  } else {
    months = months + " months";
  }

  let blog = {
    project,
    startDate,
    endDate,
    time,
    date,
    months,
    description,
    image,
    checkbox1,
    checkbox2,
    checkbox3,
    checkbox4,
  };

  dataBlog.push(blog);
  console.log(dataBlog);

  renderBlog();

  setInterval(function () {
    renderBlog();
  }, 1000);
}

function renderBlog() {
  document.getElementById("box-card").innerHTML = "";

  document.getElementById("text-project").innerHTML = "My Project";

  for (let index = 0; index < dataBlog.length; index++) {
    document.getElementById("box-card").innerHTML += `
    <div class="card p-2 shadow m-2" style="width: 18rem">
    <img style="width: 100%; height: 180px; object-fit: cover"
    src="${dataBlog[index].image}"
    alt="cat image" />
    <div>
    <a class="link-dark link-underline link-underline-opacity-0"
    href="detail.html">
    <h3>${dataBlog[index].project}</h3></a>
    <p>Duration : ${dataBlog[index].date} ${dataBlog[index].months}</p>
    </div>
    <p>${dataBlog[index].description}</p>
    <div class="mt-1">
     ${dataBlog[index].checkbox1}
     ${dataBlog[index].checkbox2}
     ${dataBlog[index].checkbox3}
     ${dataBlog[index].checkbox4}
    </div>
    <span class="mt-2">
      ${getDistance(dataBlog[index].time)}
    </span>
    <div class="row mt-1">
     <button class="col btn btn-dark ms-3 me-2">edit</button>
     <button class="col btn btn-dark ms-2 me-3">delete</button>
     </div>
    </div>
   `;
  }
}

function getDistance(time) {
  let timeNow = new Date();
  let timePost = time;

  let timeDistance = timeNow - timePost;

  let distanceSecond = Math.floor(timeDistance / 1000);
  let distanceMinutes = Math.floor(distanceSecond / 60);
  let distanceHours = Math.floor(distanceMinutes / 60);
  let distanceDays = Math.floor(distanceHours / 24);

  if (distanceSecond == 0) {
    return ``;
  } else if (distanceSecond < 60) {
    return `${distanceSecond} Second ago`;
  } else if (distanceMinutes < 60) {
    return `${distanceMinutes} Minute ago`;
  } else if (distanceHours < 60) {
    return `${distanceHours} Hous ago`;
  } else if (distanceDays < 24) {
    return `${distanceDays} Day ago`;
  }
}
