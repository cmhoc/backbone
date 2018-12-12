#used to generate the seat graphs.
using Gadfly
using Colors: RGB
using DataFrames

#sql connections to come later
function graph()
    Gadfly.with_theme(:dark) do
        #data = data()
        names = ["Libs", "NDP", "UCP", "RSP", "Ind"]
        seats = [5, 13, 5, 2, 1]
        colours = [RGB(184/255,44/255,44/255), RGB(255/255,129/255,0),
        RGB(65/255,88/255,169/255), RGB(236/255,45/255,74/255), RGB(102/255,102/255,102/255)]
        df = DataFrames.DataFrame(Parties = names, Seats = seats, Colours = colours)
        p = Gadfly.plot(df ,x=:Parties, y=:Seats, color=:Colours, Geom.bar,
        Guide.xlabel("Parties"), Guide.ylabel("Seat Counts"))
        img = SVG("seatchart.svg", 10cm, 14cm)
        draw(img, p)
    end
end

graph()
