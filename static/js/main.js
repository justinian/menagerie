var populateWildLinks = function (listElement) {
    $.ajax("/api/worlds", {
        dataType: "json",
        success: function(data, statusString, xhr) {
            for (i = 0; i < data.length; i++) {
                var item = document.createElement("li");
                var link = document.createElement("a");
                var worldName = document.createTextNode(data[i]["name"]);
                link.classList.add("dropdown-item");
                link.href = "wild.html#" + data[i]["id"]
                link.appendChild(worldName);
                item.appendChild(link);
                listElement.appendChild(item);
            }
        }
    });
}

var colorFunc = function (data, type, row, meta) {
    var c = null;
    if (data < 0) {
        c = dyes[data+55]
    } else {
        c = colors[data];
    }

    if (c) {
        return "<div class='swatch' style='background-color: "
            + c.color + ";'></div> " + c.id + ": " + c.name;
    } else {
        if (data != 0) {
            console.log("Unknown color value: " + data + " for slot " + index);
        }
        return "<div class='swatch'></div> " + data + ": Unused";
    }
}

var fixedFloat = function (data, type, row, meta) {
    return data.toFixed(0);
};

var maps = {
    "TheIsland": {"shiftx": 0.5, "shifty": 0.5, "mulx": 800000, "muly": 800000},
    "ScorchedEarth": {"shiftx": 0.5, "shifty": 0.5, "mulx": 800000, "muly": 800000},
    "Aberration": {"shiftx": 0.5, "shifty": 0.5, "mulx": 800000, "muly": 800000},
    "Extinction": {"shiftx": 0.5, "shifty": 0.5, "mulx": 800000, "muly": 800000},
    "Ragnarok": {"shiftx": 0.5, "shifty": 0.5, "mulx": 1310000, "muly": 1310000},
    "Valguero": {"shiftx": 0.5, "shifty": 0.5, "mulx": 816000, "muly": 816000},
    "CrystalIsles": {"shiftx": 0.4875, "shifty": 0.5, "mulx": 1600000, "muly": 1700000},
    "Genesis1": {"shiftx": 0.5, "shifty": 0.5, "mulx": 1050000, "muly": 1050000},
    "Genesis2": {"shiftx": 0.49655, "shifty": 0.49655, "mulx": 1450000, "muly": 1450000},
};

var drawMap = function (canvas, world, x, y) {
    var mapInfo = maps[world];
    if (mapInfo === undefined) {
        return;
    }

    var mapImage = new Image();
    mapImage.onload = function () {
        canvas.height = mapImage.height;
        canvas.width = mapImage.width;

        ctx = canvas.getContext("2d");
        ctx.drawImage(mapImage, 0, 0);

        mx = (x / mapInfo["mulx"] + mapInfo["shiftx"]) * canvas.width;
        my = (y / mapInfo["muly"] + mapInfo["shifty"]) * canvas.height;
        pointSize = canvas.width / 125;

        ctx.fillStyle = "#ff2020";
        ctx.beginPath();
        ctx.arc(mx, my, pointSize, 0, 2 * Math.PI, true);
        ctx.fill();

        canvas.style.width = "100%";
    };
    mapImage.src = "images/maps/" + world + ".webp";
};
