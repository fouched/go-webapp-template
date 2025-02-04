function resetSearch() {
    let el = document.getElementById('q');
    el.value='';
    el.focus();

    document.querySelector('.grid-scroll').scrollTop=0;
}