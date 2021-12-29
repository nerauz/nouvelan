const getScore = () => {
    axios.get('/score').then((res) => {
        if (res.status !== 200)
            return

        let highest = 0;
        let highest_key = "";

        Object.keys(res.data).forEach((key) => {
            const elem = document.getElementById(key);
            const value = parseInt(res.data[key]);

            elem.style.height = Math.floor(res.data[key] * 100 / 200) + "%"
            elem.ariaValueNow = value

            if (highest < value) {
                highest = value;
                highest_key = key;
            }

            document.getElementById(`${key}-crown`).style.visibility= "hidden";
        })

        document.getElementById(`${highest_key}-crown`).style.visibility = "visible";
    }).catch((err) => {
        console.error(err);
    })
}

document.addEventListener("DOMContentLoaded", () => {
    setInterval(getScore, 5000)
});