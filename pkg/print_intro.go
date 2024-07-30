package pkg

import "github.com/gookit/color"

const GengWithStyle = `
  	
     ██████╗ ███████╗███╗   ██╗       ██████╗ 
    ██╔════╝ ██╔════╝████╗  ██║      ██╔════╝ 
    ██║  ███╗█████╗  ██╔██╗ ██║█████╗██║  ███╗
    ██║   ██║██╔══╝  ██║╚██╗██║╚════╝██║   ██║
    ╚██████╔╝███████╗██║ ╚████║      ╚██████╔╝
     ╚═════╝ ╚══════╝╚═╝  ╚═══╝       ╚═════╝ 

    GENG: GENERATE GOLANG PROJECT
  `

// PrintIntro prints intro text with cyan color
func PrintIntro() {
	color.Cyanln(GengWithStyle)
}
