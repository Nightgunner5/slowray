SlowRAY
=======

Command line arguments
----------------------
_All arguments take a single, required parameter._

<dl>
<dt>-cpuprof</dt>
<dd>Outputs a [pprof](http://google-perftools.googlecode.com/svn/trunk/doc/cpuprofile.html) CPU profile. For debugging use only; not useful for the average user.</dd>
<dt>-cpus</dt>
<dd>The number of CPU cores to use at any given time. This defaults to the environment variable GOMAXPROCS, or the number of CPU cores if GOMAXPROCS is not set or set to 1.</dd>
<dt>-o</dt>
<dd>The (relative or absolute) path to the file that should contain the output of SlowRAY. Defaults to render.png in the current directory.</dd>
<dt>-oct</dt>
<dd>The number of octaves of perlin noise to use. Each octave of noise is scaled and layered onto the previous layers, making the output less smooth and more interesting. Defaults to 3 octaves.</dd>
<dt>-ppi</dt>
<dd>The number of points in the rendered image. One point roughly converts to a unit cube in perlin noise. Defaults to 16, resulting in a 16x16 **point** image.</dd>
<dt>-ppp</dt>
<dd>The number of pixels per point. Defaults to 8, resulting (with the default -ppi of 16) in a 128x128 pixel image.</dd>
<dt>-spp</dt>
<dd>The number of samples per pixel, in both x and y directions. This allows multisampled antialiasing. Defaults to 4, which means 16 samples per pixel.</dd>
</dl>
