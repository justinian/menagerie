var colors = [
    null,
    {"id": 1, "name": "Red", "color": "#ff0000"},
    {"id": 2, "name": "Blue", "color": "#0000ff"},
    {"id": 3, "name": "Green", "color": "#00ff00"},
    {"id": 4, "name": "Yellow", "color": "#ffff00"},
    {"id": 5, "name": "Cyan", "color": "#00ffff"},
    {"id": 6, "name": "Magenta", "color": "#ff00ff"},
    {"id": 7, "name": "Light Green", "color": "#c0ffba"},
    {"id": 8, "name": "Light Grey", "color": "#c8caca"},
    {"id": 9, "name": "Light Brown", "color": "#786759"},
    {"id": 10, "name": "Light Orange", "color": "#ffb46c"},
    {"id": 11, "name": "Light Yellow", "color": "#fffa8a"},
    {"id": 12, "name": "Light Red", "color": "#ff756c"},
    {"id": 13, "name": "Dark Grey", "color": "#7b7b7b"},
    {"id": 14, "name": "Black", "color": "#3b3b3b"},
    {"id": 15, "name": "Brown", "color": "#593a2a"},
    {"id": 16, "name": "Dark Green", "color": "#224900"},
    {"id": 17, "name": "Dark Red", "color": "#812118"},
    {"id": 18, "name": "White", "color": "#ffffff"},
    {"id": 19, "name": "Dino Light Red", "color": "#ffa8a8"},
    {"id": 20, "name": "Dino Dark Red", "color": "#592b2b"},
    {"id": 21, "name": "Dino Light Orange", "color": "#ffb694"},
    {"id": 22, "name": "Dino Dark Orange", "color": "#88532f"},
    {"id": 23, "name": "Dino Light Yellow", "color": "#cacaa0"},
    {"id": 24, "name": "Dino Dark Yellow", "color": "#94946c"},
    {"id": 25, "name": "Dino Light Green", "color": "#e0ffe0"},
    {"id": 26, "name": "Dino Medium Green", "color": "#799479"},
    {"id": 27, "name": "Dino Dark Green", "color": "#224122"},
    {"id": 28, "name": "Dino Light Blue", "color": "#d9e0ff"},
    {"id": 29, "name": "Dino Dark Blue", "color": "#394263"},
    {"id": 30, "name": "Dino Light Purple", "color": "#e4d9ff"},
    {"id": 31, "name": "Dino Dark Purple", "color": "#403459"},
    {"id": 32, "name": "Dino Light Brown", "color": "#ffe0ba"},
    {"id": 33, "name": "Dino Medium Brown", "color": "#948575"},
    {"id": 34, "name": "Dino Dark Brown", "color": "#594e41"},
    {"id": 35, "name": "Dino Darker Grey", "color": "#595959"},
    {"id": 36, "name": "Dino Albino", "color": "#ffffff"},
    {"id": 37, "name": "BigFoot0", "color": "#b79683"},
    {"id": 38, "name": "BigFoot4", "color": "#eadad5"},
    {"id": 39, "name": "BigFoot5", "color": "#d0a794"},
    {"id": 40, "name": "WolfFur", "color": "#c3b39f"},
    {"id": 41, "name": "DarkWolfFur", "color": "#887666"},
    {"id": 42, "name": "DragonBase0", "color": "#a0664b"},
    {"id": 43, "name": "DragonBase1", "color": "#cb7956"},
    {"id": 44, "name": "DragonFire", "color": "#bc4f00"},
    {"id": 45, "name": "DragonGreen0", "color": "#79846c"},
    {"id": 46, "name": "DragonGreen1", "color": "#909c79"},
    {"id": 47, "name": "DragonGreen2", "color": "#a5a48b"},
    {"id": 48, "name": "DragonGreen3", "color": "#74939c"},
    {"id": 49, "name": "WyvernPurple0", "color": "#787496"},
    {"id": 50, "name": "WyvernPurple1", "color": "#b0a2c0"},
    {"id": 51, "name": "WyvernBlue0", "color": "#6281a7"},
    {"id": 52, "name": "WyvernBlue1", "color": "#485c75"},
    {"id": 53, "name": "Dino Medium Blue", "color": "#5fa4ea"},
    {"id": 54, "name": "Dino Deep Blue", "color": "#4568d4"},
    {"id": 55, "name": "NearWhite", "color": "#ededed"},
    {"id": 56, "name": "NearBlack", "color": "#515151"},
    {"id": 57, "name": "DarkTurquoise", "color": "#184546"},
    {"id": 58, "name": "MediumTurquoise", "color": "#007060"},
    {"id": 59, "name": "Turquoise", "color": "#00c5ab"},
    {"id": 60, "name": "GreenSlate", "color": "#40594c"},
    {"id": 61, "name": "Sage", "color": "#3e4f40"},
    {"id": 62, "name": "DarkWarmGray", "color": "#3b3938"},
    {"id": 63, "name": "MediumWarmGray", "color": "#585554"},
    {"id": 64, "name": "LightWarmGray", "color": "#9b9290"},
    {"id": 65, "name": "DarkCement", "color": "#525b56"},
    {"id": 66, "name": "LightCement", "color": "#8aa196"},
    {"id": 67, "name": "LightPink", "color": "#e8b0ff"},
    {"id": 68, "name": "DeepPink", "color": "#ff119a"},
    {"id": 69, "name": "DarkViolet", "color": "#730046"},
    {"id": 70, "name": "DarkMagenta", "color": "#b70042"},
    {"id": 71, "name": "BurntSienna", "color": "#7e331e"},
    {"id": 72, "name": "MediumAutumn", "color": "#a93000"},
    {"id": 73, "name": "Vermillion", "color": "#ef3100"},
    {"id": 74, "name": "Coral", "color": "#ff5834"},
    {"id": 75, "name": "Orange", "color": "#ff7f00"},
    {"id": 76, "name": "Peach", "color": "#ffa73a"},
    {"id": 77, "name": "LightAutumn", "color": "#ae7000"},
    {"id": 78, "name": "Mustard", "color": "#949427"},
    {"id": 79, "name": "ActualBlack", "color": "#171717"},
    {"id": 80, "name": "MidnightBlue", "color": "#191d36"},
    {"id": 81, "name": "DarkBlue", "color": "#152b3a"},
    {"id": 82, "name": "BlackSands", "color": "#302531"},
    {"id": 83, "name": "LemonLime", "color": "#a8ff44"},
    {"id": 84, "name": "Mint", "color": "#38e985"},
    {"id": 85, "name": "Jade", "color": "#008840"},
    {"id": 86, "name": "PineGreen", "color": "#0f552e"},
    {"id": 87, "name": "SpruceGreen", "color": "#005b45"},
    {"id": 88, "name": "LeafGreen", "color": "#5b9725"},
    {"id": 89, "name": "DarkLavender", "color": "#5e275f"},
    {"id": 90, "name": "MediumLavender", "color": "#853587"},
    {"id": 91, "name": "Lavender", "color": "#bd77be"},
    {"id": 92, "name": "DarkTeal", "color": "#0e404a"},
    {"id": 93, "name": "MediumTeal", "color": "#105563"},
    {"id": 94, "name": "Teal", "color": "#14849c"},
    {"id": 95, "name": "PowderBlue", "color": "#82a7ff"},
    {"id": 96, "name": "Glacial", "color": "#aceaff"},
    {"id": 97, "name": "Cammo", "color": "#505118"},
    {"id": 98, "name": "DryMoss", "color": "#766e3f"},
    {"id": 99, "name": "Custard", "color": "#c0bd5e"},
    {"id": 100, "name": "Cream", "color": "#f4ffc0"},
]

var dyes = [
    {"id": 201, "name": "Black Dye", "color": "#1f1f1f"},
    {"id": 202, "name": "Blue Dye", "color": "#0000ff"},
    {"id": 203, "name": "Brown Dye", "color": "#756147"},
    {"id": 204, "name": "Cyan Dye", "color": "#00ffff"},
    {"id": 205, "name": "Forest Dye", "color": "#006c00"},
    {"id": 206, "name": "Green Dye", "color": "#00ff00"},
    {"id": 207, "name": "Unused Purple Dye", "color": "#6c00ba"},
    {"id": 208, "name": "Orange Dye", "color": "#ff8800"},
    {"id": 209, "name": "Parchment Dye", "color": "#ffffba"},
    {"id": 210, "name": "Pink Dye", "color": "#ff7be1"},
    {"id": 211, "name": "Purple Dye", "color": "#7b00e0"},
    {"id": 212, "name": "Red Dye", "color": "#ff0000"},
    {"id": 213, "name": "Royalty Dye", "color": "#7b00a8"},
    {"id": 214, "name": "Silver Dye", "color": "#e0e0e0"},
    {"id": 215, "name": "Sky Dye", "color": "#bad4ff"},
    {"id": 216, "name": "Tan Dye", "color": "#ffed82"},
    {"id": 217, "name": "Tangerine Dye", "color": "#ad652c"},
    {"id": 218, "name": "White Dye", "color": "#fefefe"},
    {"id": 219, "name": "Yellow Dye", "color": "#ffff00"},
    {"id": 220, "name": "Magenta Dye", "color": "#e71fd9"},
    {"id": 221, "name": "Brick Dye", "color": "#94341f"},
    {"id": 222, "name": "Cantaloupe Dye", "color": "#ff9a00"},
    {"id": 223, "name": "Mud Dye", "color": "#473b2b"},
    {"id": 224, "name": "Navy Dye", "color": "#34346c"},
    {"id": 225, "name": "Olive Dye", "color": "#baba59"},
    {"id": 226, "name": "Slate Dye", "color": "#595959"},
]