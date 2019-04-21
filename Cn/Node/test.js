function qsort(arr) {
    if (!arr) return [];
    if (arr.length <= 0) return [];

    qsortCore(arr, 0, arr.length - 1);
}

function qsortCore(arr,  start, end) {
    let index = partition(arr, start, end);
    if (index > start) {
        qsortCore(arr, start, index - 1);
    }

    if (index < end) {
        qsortCore(arr, index + 1, end);
    }
}

function partition(arr, start, end) {
    const len = end - start + 1;
    const theIndex = Math.floor(Math.random() * len);

    [arr[end], arr[start + theIndex]] = [arr[start + theIndex], arr[end]];

    let theVal = arr[end];

    let cur = start - 1;
    for(let i = 0; i < len; i++) {
        if (arr[start + i] < theVal) {
            cur++;
            [arr[cur], arr[start + i]] = [arr[start + i], arr[cur]];
        }
    }
    
    cur++;
    [arr[cur], arr[end]] = [arr[end], arr[cur]];

    return cur;
}

let a = [0, 3, 88, 88, 9, 9, 5, -22, -33, 66, 0];

qsort(a);

console.log(a);