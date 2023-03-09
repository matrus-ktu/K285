document.addEventListener("DOMContentLoaded", function() {
    window.addEventListener("scroll", function() {
        let scrollHeight = window.scrollY;
        let backToTop = document.querySelector(".back-to-top");
        if (scrollHeight > 200) {
            backToTop.style.display = "block";
        } else {
            backToTop.style.display = "none";
        }
    });

    let backToTopButton = document.querySelector(".back-to-top");
    backToTopButton.addEventListener("click", function() {
        window.scrollTo({
            top: 0,
            behavior: "smooth"
        });
    });
});