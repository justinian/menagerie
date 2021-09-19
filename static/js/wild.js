
var showStats = function () {
    var showBase = $( '#showBase' )[0].checked;
    var showColor = $( '#showColor' )[0].checked;

    table.columns([6, 7, 8, 9, 10, 11]).visible(showColor);
    table.columns([12, 13, 14, 15, 16, 17, 18, 19]).visible(showBase);
};

var columns = [
    {"data": "world", "title": "World", "visible": false},
    {"data": "class_name", "title": "Class"},
    {"data": "levels_wild", "title": "Level"},

    {"data": "x", "visible": false},
    {"data": "y", "visible": false},
    {"data": "z", "visible": false},

    {"data": "color0", "render": colorFunc, "title": "C0", "searchBuilderTitle": "Color 0", "visible": false},
    {"data": "color1", "render": colorFunc, "title": "C1", "searchBuilderTitle": "Color 1", "visible": false},
    {"data": "color2", "render": colorFunc, "title": "C2", "searchBuilderTitle": "Color 2", "visible": false},
    {"data": "color3", "render": colorFunc, "title": "C3", "searchBuilderTitle": "Color 3", "visible": false},
    {"data": "color4", "render": colorFunc, "title": "C4", "searchBuilderTitle": "Color 4", "visible": false},
    {"data": "color5", "render": colorFunc, "title": "C5", "searchBuilderTitle": "Color 5", "visible": false},

    {"data": "health_wild", "title": "H", "searchBuilderTitle": "Base Health"},
    {"data": "stamina_wild", "title": "St", "searchBuilderTitle": "Base Stamina"},
    {"data": "torpor_wild", "title": "T", "searchBuilderTitle": "Base Torpor"},
    {"data": "oxygen_wild", "title": "O", "searchBuilderTitle": "Base Oxygen"},
    {"data": "food_wild", "title": "F", "searchBuilderTitle": "Base Food"},
    {"data": "weight_wild", "title": "W", "searchBuilderTitle": "Base Weight"},
    {"data": "melee_wild", "title": "M", "searchBuilderTitle": "Base Melee"},
    {"data": "speed_wild", "title": "Sp", "searchBuilderTitle": "Base Speed"}
];

var tableOptions = {
    "ajax": {"dataSrc":""},
    "columns": columns,

    "dom": "rtpil",
    "pageLength": 50,
    "scrollX": true,
    "processing": true,

    "language": {
        "searchBuilder": {
            "title": ""
        },
    },

    "select": {
        "info": false,
        "style": "single",
    },

    "searchBuilder": {
        "columns": [0, 1, 2,
                    6, 7, 8, 9, 10, 11,
                    12, 13, 14, 15, 16, 17, 18, 19]
    }
};

var setTableWorld = function (table) {
    var id = window.location.hash.substr(1);
    if (id.length > 0) {
        table.rows({selected: true}).deselect();
        url = "/api/worlds/" + id + "/dinos";
        table.ajax.url(url).load();
    }
};
