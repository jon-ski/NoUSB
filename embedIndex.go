package main

const bundleIndex = `<!doctype html>
<html>

<head>
  <title>nousb</title>
  <link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:300,300italic,700,700italic">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>

<body>
  <style type="text/css">
    .logo {
      display: block;
      margin: 2em 0;
      height: 4em;
      fill: #000000;
    }
  </style>
  <div class="container">
    <svg class="logo" viewBox="0 0 324 76" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink"
      xml:space="preserve" xmlns:serif="http://www.serif.com/" style="fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:1.41421;">
      <g>
        <path d="M0,68.509c0,-3.308 0.106,-6.615 0.317,-9.923c0.211,-3.308 0.475,-6.598 0.791,-9.87c0.317,-3.272 0.581,-6.527 0.792,-9.764c0.211,-3.308 0.317,-6.58 0.317,-9.817c0,-1.83 -0.194,-3.589 -0.581,-5.278c-0.387,-1.689 -0.58,-3.413 -0.58,-5.173l0,-0.528c0.14,-0.07 0.404,-0.158 0.791,-0.264c0.387,-0.105 0.792,-0.193 1.214,-0.263c0.422,-0.071 0.845,-0.176 1.267,-0.317c0.352,-0.141 0.598,-0.211 0.739,-0.211l1.055,0c0.563,2.533 1.32,5.559 2.27,9.078c0.95,3.519 2.094,7.108 3.431,10.767c1.337,3.66 2.92,7.301 4.75,10.926c1.83,3.624 3.941,6.861 6.334,9.711c2.392,2.85 5.066,5.137 8.022,6.862c2.956,1.724 6.263,2.551 9.923,2.48c2.181,0 4.099,-0.369 5.753,-1.108c1.654,-0.739 3.008,-1.76 4.064,-3.061c1.056,-1.302 1.847,-2.833 2.375,-4.592c0.528,-1.76 0.792,-3.695 0.792,-5.806c0,-2.745 -0.546,-5.296 -1.636,-7.653c-1.091,-2.358 -2.27,-4.627 -3.537,-6.809c-1.266,-2.181 -2.445,-4.363 -3.536,-6.545c-1.091,-2.181 -1.636,-4.539 -1.636,-7.072c0,-0.985 0.105,-1.759 0.317,-2.322c0.211,-0.563 0.809,-0.845 1.794,-0.845c1.126,0 2.27,0.633 3.431,1.9c1.161,1.267 2.269,2.885 3.325,4.856c1.056,1.97 2.041,4.17 2.956,6.597c0.914,2.428 1.724,4.839 2.427,7.231c0.704,2.393 1.25,4.61 1.637,6.651c0.387,2.04 0.58,3.659 0.58,4.855c0,3.026 -0.492,5.789 -1.478,8.287c-0.985,2.498 -2.375,4.68 -4.169,6.545c-1.795,1.865 -3.941,3.307 -6.439,4.328c-2.499,1.02 -5.261,1.53 -8.287,1.53c-3.307,0 -6.281,-0.633 -8.92,-1.9c-2.639,-1.267 -4.996,-2.903 -7.072,-4.908c-2.076,-2.006 -3.906,-4.135 -5.489,-6.387c-1.584,-2.252 -2.939,-4.381 -4.064,-6.386c-1.126,-2.006 -2.094,-3.642 -2.903,-4.909c-0.81,-1.267 -1.461,-1.9 -1.953,-1.9c-0.915,0 -1.584,0.317 -2.006,0.95c-0.422,0.634 -0.721,1.372 -0.897,2.217c-0.176,0.844 -0.264,1.742 -0.264,2.692c0,0.95 0,1.671 0,2.164c0,2.252 0.211,4.345 0.633,6.281c0.423,1.935 0.634,3.993 0.634,6.175c0,0.492 -0.053,1.126 -0.159,1.9c-0.105,0.774 -0.281,1.548 -0.527,2.322c-0.247,0.774 -0.616,1.443 -1.109,2.006c-0.492,0.563 -1.126,0.844 -1.9,0.844c-0.704,0 -1.284,-0.264 -1.742,-0.791c-0.457,-0.528 -0.827,-1.126 -1.108,-1.795c-0.282,-0.668 -0.475,-1.372 -0.581,-2.111c-0.105,-0.739 -0.158,-1.355 -0.158,-1.847Z"
          style="fill-rule:nonzero;" />
        <path d="M64.814,53.414c0,-3.871 0.844,-7.249 2.533,-10.134c1.689,-2.886 3.906,-5.313 6.651,-7.284c2.744,-1.97 5.823,-3.431 9.236,-4.381c3.413,-0.95 6.844,-1.425 10.292,-1.425c1.83,0 3.537,0.335 5.12,1.003c1.583,0.669 3.097,1.583 4.539,2.745c1.443,1.161 2.727,2.498 3.853,4.011c1.126,1.513 2.094,3.079 2.903,4.697c0.809,1.619 1.425,3.255 1.847,4.909c0.423,1.654 0.634,3.255 0.634,4.803c0,3.589 -0.704,6.791 -2.112,9.606c-1.407,2.815 -3.307,5.19 -5.7,7.125c-2.392,1.935 -5.102,3.413 -8.128,4.434c-3.026,1.02 -6.122,1.53 -9.289,1.53c-3.237,0 -6.211,-0.51 -8.92,-1.53c-2.709,-1.021 -5.067,-2.446 -7.073,-4.275c-2.005,-1.83 -3.571,-4.1 -4.697,-6.809c-1.126,-2.709 -1.689,-5.718 -1.689,-9.025Zm5.172,-0.845c0,2.182 0.405,4.275 1.214,6.281c0.81,2.006 1.936,3.747 3.378,5.225c1.443,1.478 3.15,2.657 5.12,3.536c1.97,0.88 4.082,1.32 6.334,1.32c1.196,0.141 2.656,-0.053 4.38,-0.581c1.725,-0.528 3.449,-1.302 5.173,-2.322c1.724,-1.02 3.378,-2.252 4.961,-3.695c1.584,-1.442 2.903,-3.026 3.959,-4.75c1.055,-1.724 1.706,-3.519 1.953,-5.383c0.246,-1.865 -0.071,-3.73 -0.95,-5.595c-0.88,-1.865 -2.481,-3.66 -4.803,-5.384c-2.323,-1.724 -5.525,-3.325 -9.606,-4.803c-2.323,0 -4.733,0.335 -7.231,1.003c-2.499,0.669 -4.768,1.671 -6.809,3.009c-2.041,1.337 -3.73,3.026 -5.067,5.066c-1.337,2.041 -2.006,4.399 -2.006,7.073Z"
          style="fill-rule:nonzero;" />
        <path d="M118.439,33.779c0,-0.633 0.017,-1.495 0.052,-2.586c0.036,-1.091 0.106,-2.287 0.212,-3.589c0.105,-1.302 0.299,-2.639 0.58,-4.011c0.282,-1.373 0.704,-2.604 1.267,-3.695c0.563,-1.091 1.267,-1.988 2.111,-2.692c0.845,-0.703 1.865,-1.055 3.061,-1.055c0.774,0 1.373,0.211 1.795,0.633c0.422,0.422 0.668,0.933 0.739,1.531c0.07,0.598 -0.106,1.249 -0.528,1.953c-0.422,0.703 -1.091,1.337 -2.006,1.9c-0.633,0 -1.09,0.475 -1.372,1.425c-0.281,0.95 -0.493,1.988 -0.633,3.114c-0.141,1.126 -0.211,2.217 -0.211,3.272c0,1.056 0,1.689 0,1.9c0,1.83 0.07,3.941 0.211,6.334c0.14,2.393 0.422,4.873 0.844,7.442c0.422,2.568 1.038,5.119 1.847,7.653c0.81,2.533 1.865,4.838 3.167,6.914c1.302,2.076 2.868,3.818 4.698,5.225c1.829,1.408 4.046,2.182 6.65,2.323c1.97,0.14 4.258,-0.211 6.861,-1.056c2.604,-0.844 5.032,-2.058 7.284,-3.642c2.252,-1.583 4.17,-3.413 5.753,-5.489c1.583,-2.076 2.375,-4.24 2.375,-6.492l0,-0.422c-0.492,-2.956 -1.003,-5.419 -1.53,-7.389c-0.528,-1.971 -0.986,-3.765 -1.373,-5.384c-0.387,-1.618 -0.703,-3.272 -0.95,-4.961c-0.246,-1.689 -0.369,-3.66 -0.369,-5.912c0.211,-3.237 0.58,-5.647 1.108,-7.23c0.528,-1.584 1.073,-2.464 1.636,-2.639c0.563,-0.176 1.162,0.228 1.795,1.213c0.633,0.986 1.196,2.411 1.689,4.276c0.492,1.865 0.844,4.081 1.055,6.65c0.212,2.569 0.282,5.296 0.212,8.181c0.07,2.111 0.475,4.433 1.214,6.967c0.738,2.533 1.565,5.067 2.48,7.6c0.915,2.534 1.847,4.997 2.798,7.389c0.95,2.393 1.671,4.487 2.164,6.281c0.492,1.795 0.633,3.273 0.422,4.434c-0.211,1.161 -0.95,1.742 -2.217,1.742c-1.407,0 -2.463,-0.423 -3.167,-1.267c-0.704,-0.845 -1.302,-1.777 -1.794,-2.798c-0.493,-1.02 -1.021,-1.952 -1.584,-2.797c-0.563,-0.844 -1.266,-1.267 -2.111,-1.267c-0.704,0 -1.548,0.264 -2.533,0.792c-0.986,0.528 -2.094,1.214 -3.326,2.059c-1.231,0.844 -2.568,1.759 -4.011,2.744c-1.442,0.985 -2.938,1.9 -4.486,2.745c-1.548,0.844 -3.132,1.548 -4.75,2.111c-1.619,0.563 -3.273,0.844 -4.962,0.844c-3.026,0 -5.665,-0.721 -7.917,-2.164c-2.252,-1.442 -4.187,-3.325 -5.806,-5.647c-1.618,-2.322 -2.973,-4.979 -4.064,-7.97c-1.09,-2.991 -1.952,-6.017 -2.586,-9.078c-0.633,-3.061 -1.091,-6.017 -1.372,-8.867c-0.282,-2.85 -0.422,-5.366 -0.422,-7.548Z"
          style="fill-rule:nonzero;" />
        <path d="M179.347,56.158c0,-0.704 0.141,-1.46 0.422,-2.269c0.282,-0.81 0.669,-1.584 1.161,-2.323c0.493,-0.739 1.091,-1.354 1.795,-1.847c0.704,-0.493 1.443,-0.739 2.217,-0.739c0,4.574 0.809,8.181 2.428,10.82c1.618,2.639 3.852,4.662 6.703,6.07c2.85,1.407 6.104,2.357 9.764,2.85c3.659,0.492 7.6,0.844 11.823,1.055c1.407,-0.07 3.254,-0.123 5.542,-0.158c2.287,-0.035 4.486,-0.281 6.597,-0.739c2.111,-0.457 3.906,-1.214 5.384,-2.269c1.478,-1.056 2.181,-2.675 2.111,-4.856c0,-0.985 -0.035,-1.865 -0.106,-2.639c-0.07,-0.774 -0.228,-1.46 -0.475,-2.059c-0.246,-0.598 -0.668,-1.178 -1.266,-1.741c-0.599,-0.563 -1.425,-1.126 -2.481,-1.689l-31.352,-17.101c-1.125,-0.634 -2.216,-1.425 -3.272,-2.375c-1.055,-0.95 -1.935,-2.006 -2.639,-3.167c-0.704,-1.161 -1.267,-2.393 -1.689,-3.695c-0.422,-1.302 -0.633,-2.586 -0.633,-3.853c0,-2.674 0.721,-4.855 2.164,-6.544c1.442,-1.689 3.237,-3.044 5.383,-4.064c2.147,-1.021 4.469,-1.742 6.967,-2.164c2.499,-0.423 4.803,-0.634 6.915,-0.634c0.985,0 2.251,0.088 3.8,0.264c1.548,0.176 3.202,0.44 4.961,0.792c1.759,0.352 3.501,0.809 5.225,1.372c1.724,0.563 3.255,1.284 4.592,2.164c1.337,0.88 2.446,1.9 3.325,3.061c0.88,1.162 1.32,2.516 1.32,4.064c-0.985,0.563 -1.988,0.739 -3.009,0.528c-1.02,-0.211 -2.058,-0.598 -3.114,-1.161c-1.055,-0.563 -2.146,-1.196 -3.272,-1.9c-1.126,-0.704 -2.287,-1.407 -3.484,-2.111c-1.196,-0.704 -2.463,-1.267 -3.8,-1.689c-1.337,-0.422 -2.744,-0.563 -4.222,-0.422c-1.337,0 -3.061,0.07 -5.173,0.211c-2.111,0.14 -4.169,0.475 -6.175,1.003c-2.006,0.527 -3.73,1.337 -5.172,2.427c-1.443,1.091 -2.164,2.622 -2.164,4.592c0,1.9 0.668,3.642 2.005,5.225c1.337,1.584 3.15,3.097 5.437,4.54c2.287,1.442 4.855,2.815 7.706,4.116c2.85,1.302 5.735,2.639 8.656,4.012c2.92,1.372 5.77,2.797 8.55,4.275c2.78,1.478 5.208,3.096 7.284,4.856c2.076,1.759 3.694,3.642 4.855,5.647c1.162,2.006 1.566,4.275 1.214,6.809c-0.352,2.533 -1.442,4.715 -3.272,6.545c-1.83,1.829 -4.152,3.254 -6.967,4.275c-2.815,1.02 -5.947,1.689 -9.395,2.005c-3.448,0.317 -7.002,0.299 -10.662,-0.052c-3.659,-0.352 -7.283,-1.021 -10.872,-2.006c-3.589,-0.985 -6.862,-2.305 -9.817,-3.959c-2.956,-1.653 -5.489,-3.589 -7.601,-5.805c-2.111,-2.217 -3.518,-4.733 -4.222,-7.548Z"
          style="fill-rule:nonzero;" />
        <path d="M277.624,71.887c-0.141,-0.071 -0.687,-0.352 -1.637,-0.845c-0.95,-0.493 -1.935,-1.091 -2.955,-1.794c-1.021,-0.704 -1.918,-1.443 -2.692,-2.217c-0.774,-0.774 -1.091,-1.513 -0.95,-2.217c0.281,-0.774 0.721,-1.319 1.319,-1.636c0.599,-0.317 1.197,-0.616 1.795,-0.897c0.598,-0.282 1.108,-0.669 1.531,-1.162c0.422,-0.492 0.598,-1.302 0.527,-2.427c-1.196,-8.375 -2.287,-15.782 -3.272,-22.221c-0.985,-6.439 -2.023,-11.84 -3.114,-16.203c-1.091,-4.364 -2.287,-7.671 -3.589,-9.923c-1.302,-2.252 -2.903,-3.378 -4.803,-3.378c-0.071,0 -0.106,0 -0.106,0l-0.211,0l-0.211,0c-0.563,0.352 -1.319,0.721 -2.269,1.108c-0.951,0.387 -1.971,0.774 -3.062,1.162c-1.091,0.387 -2.164,0.738 -3.219,1.055c-1.056,0.317 -1.9,0.581 -2.534,0.792c-0.141,0 -0.316,0.035 -0.528,0.105c-0.14,0 -0.281,0 -0.422,0l0,-3.166c0,-0.704 0.634,-1.373 1.9,-2.006c1.267,-0.633 2.868,-1.249 4.803,-1.847c1.936,-0.599 4.064,-1.162 6.387,-1.689c2.322,-0.528 4.556,-0.968 6.703,-1.32c2.146,-0.352 4.064,-0.633 5.753,-0.844c1.689,-0.211 2.815,-0.317 3.378,-0.317c1.266,0 2.779,0.141 4.539,0.422c1.759,0.282 3.606,0.704 5.542,1.267c1.935,0.563 3.835,1.284 5.7,2.164c1.865,0.88 3.536,1.882 5.014,3.008c1.478,1.126 2.674,2.393 3.589,3.801c0.915,1.337 1.372,2.85 1.372,4.539c0,1.829 -0.369,3.518 -1.108,5.067c-0.739,1.548 -1.583,2.938 -2.533,4.169c-0.95,1.232 -1.795,2.323 -2.534,3.273c-0.739,0.95 -1.108,1.847 -1.108,2.691c0,0.141 0.686,0.546 2.058,1.214c1.372,0.669 3.132,1.601 5.278,2.798c2.147,1.196 4.451,2.674 6.914,4.433c2.463,1.76 4.768,3.712 6.915,5.859c2.146,2.146 3.923,4.468 5.33,6.967c1.408,2.498 2.112,5.155 2.112,7.97c0,2.955 -0.651,5.418 -1.953,7.389c-1.302,1.97 -2.974,3.536 -5.014,4.697c-2.041,1.161 -4.311,2.006 -6.809,2.534c-2.498,0.527 -4.909,0.791 -7.231,0.791c-2.322,0 -4.381,-0.052 -6.175,-0.158c-1.795,-0.106 -3.589,-0.281 -5.384,-0.528c-1.794,-0.246 -3.712,-0.563 -5.753,-0.95c-2.041,-0.387 -4.469,-0.897 -7.283,-1.53Zm3.061,-12.562c-0.915,1.407 -1.197,2.639 -0.845,3.694c0.352,1.056 1.126,1.971 2.323,2.745c1.196,0.774 2.639,1.39 4.328,1.847c1.689,0.458 3.378,0.81 5.067,1.056c1.688,0.246 3.237,0.369 4.644,0.369c1.408,0 2.463,-0.035 3.167,-0.105c2.111,0 4.187,-0.088 6.228,-0.264c2.041,-0.176 3.835,-0.633 5.384,-1.372c1.548,-0.739 2.779,-1.812 3.694,-3.22c0.915,-1.407 1.372,-3.343 1.372,-5.806c0,-3.448 -0.897,-6.316 -2.691,-8.603c-1.795,-2.287 -4.047,-4.24 -6.756,-5.858c-2.71,-1.619 -5.648,-2.974 -8.814,-4.065c-3.167,-1.09 -6.123,-2.111 -8.867,-3.061c-2.745,-0.95 -4.997,-1.953 -6.756,-3.008c-1.76,-1.056 -2.639,-2.323 -2.639,-3.8c0,-0.775 0.404,-1.39 1.214,-1.848c0.809,-0.457 1.829,-0.862 3.061,-1.214c1.231,-0.352 2.551,-0.721 3.958,-1.108c1.408,-0.387 2.745,-0.95 4.012,-1.689c1.266,-0.739 2.304,-1.671 3.114,-2.797c0.809,-1.126 1.214,-2.639 1.214,-4.539c0,-1.549 -0.423,-2.921 -1.267,-4.117c-0.845,-1.197 -1.918,-2.217 -3.22,-3.062c-1.302,-0.844 -2.709,-1.477 -4.222,-1.9c-1.513,-0.422 -2.938,-0.633 -4.275,-0.633c-1.971,0 -4.012,0.211 -6.123,0.633c-2.111,0.423 -3.941,1.267 -5.489,2.534c0.845,3.589 1.795,7.53 2.85,11.823c1.056,4.292 2.059,8.638 3.009,13.036c0.95,4.399 1.741,8.709 2.375,12.931c0.633,4.223 0.95,8.023 0.95,11.401Z"
          style="fill-rule:nonzero;" />
      </g>
    </svg>
    <a class="button" href="/download/">View Files</a>
    <a class="button button-outline" href="/downloadself/">Get nousb</a>
    <hr>
    <h3>Serve Directory</h3>
    <pre><code>
      nousb
      // Custom port (default is 8080)
      nousb -p=80
    </code></pre>
    <h3>Download Directory</h3>
    <pre><code>
      nousb -d -a=192.168.1.50:8080
    </code></pre>
    <p>NoUSB operates in the directory that it was started in.</p>
  </div>
  <style type="text/css">
    /*! normalize.css v8.0.0 | MIT License | github.com/necolas/normalize.css */

    /* Document
       ========================================================================== */

    /**
     * 1. Correct the line height in all browsers.
     * 2. Prevent adjustments of font size after orientation changes in iOS.
     */

    html {
      line-height: 1.15;
      /* 1 */
      -webkit-text-size-adjust: 100%;
      /* 2 */
    }

    /* Sections
       ========================================================================== */

    /**
     * Remove the margin in all browsers.
     */

    body {
      margin: 0;
    }

    /**
     * Correct the font size and margin on "h1" elements within "section" and
     * "article" contexts in Chrome, Firefox, and Safari.
     */

    h1 {
      font-size: 2em;
      margin: 0.67em 0;
    }

    /* Grouping content
       ========================================================================== */

    /**
     * 1. Add the correct box sizing in Firefox.
     * 2. Show the overflow in Edge and IE.
     */

    hr {
      box-sizing: content-box;
      /* 1 */
      height: 0;
      /* 1 */
      overflow: visible;
      /* 2 */
    }

    /**
     * 1. Correct the inheritance and scaling of font size in all browsers.
     * 2. Correct the odd "em" font sizing in all browsers.
     */

    pre {
      font-family: monospace, monospace;
      /* 1 */
      font-size: 1em;
      /* 2 */
    }

    /* Text-level semantics
       ========================================================================== */

    /**
     * Remove the gray background on active links in IE 10.
     */

    a {
      background-color: transparent;
    }

    /**
     * 1. Remove the bottom border in Chrome 57-
     * 2. Add the correct text decoration in Chrome, Edge, IE, Opera, and Safari.
     */

    abbr[title] {
      border-bottom: none;
      /* 1 */
      text-decoration: underline;
      /* 2 */
      text-decoration: underline dotted;
      /* 2 */
    }

    /**
     * Add the correct font weight in Chrome, Edge, and Safari.
     */

    b,
    strong {
      font-weight: bolder;
    }

    /**
     * 1. Correct the inheritance and scaling of font size in all browsers.
     * 2. Correct the odd "em" font sizing in all browsers.
     */

    code,
    kbd,
    samp {
      font-family: monospace, monospace;
      /* 1 */
      font-size: 1em;
      /* 2 */
    }

    /**
     * Add the correct font size in all browsers.
     */

    small {
      font-size: 80%;
    }

    /**
     * Prevent "sub" and "sup" elements from affecting the line height in
     * all browsers.
     */

    sub,
    sup {
      font-size: 75%;
      line-height: 0;
      position: relative;
      vertical-align: baseline;
    }

    sub {
      bottom: -0.25em;
    }

    sup {
      top: -0.5em;
    }

    /* Embedded content
       ========================================================================== */

    /**
     * Remove the border on images inside links in IE 10.
     */

    img {
      border-style: none;
    }

    /* Forms
       ========================================================================== */

    /**
     * 1. Change the font styles in all browsers.
     * 2. Remove the margin in Firefox and Safari.
     */

    button,
    input,
    optgroup,
    select,
    textarea {
      font-family: inherit;
      /* 1 */
      font-size: 100%;
      /* 1 */
      line-height: 1.15;
      /* 1 */
      margin: 0;
      /* 2 */
    }

    /**
     * Show the overflow in IE.
     * 1. Show the overflow in Edge.
     */

    button,
    input {
      /* 1 */
      overflow: visible;
    }

    /**
     * Remove the inheritance of text transform in Edge, Firefox, and IE.
     * 1. Remove the inheritance of text transform in Firefox.
     */

    button,
    select {
      /* 1 */
      text-transform: none;
    }

    /**
     * Correct the inability to style clickable types in iOS and Safari.
     */

    button,
    [type="button"],
    [type="reset"],
    [type="submit"] {
      -webkit-appearance: button;
    }

    /**
     * Remove the inner border and padding in Firefox.
     */

    button::-moz-focus-inner,
    [type="button"]::-moz-focus-inner,
    [type="reset"]::-moz-focus-inner,
    [type="submit"]::-moz-focus-inner {
      border-style: none;
      padding: 0;
    }

    /**
     * Restore the focus styles unset by the previous rule.
     */

    button:-moz-focusring,
    [type="button"]:-moz-focusring,
    [type="reset"]:-moz-focusring,
    [type="submit"]:-moz-focusring {
      outline: 1px dotted ButtonText;
    }

    /**
     * Correct the padding in Firefox.
     */

    fieldset {
      padding: 0.35em 0.75em 0.625em;
    }

    /**
     * 1. Correct the text wrapping in Edge and IE.
     * 2. Correct the color inheritance from "fieldset" elements in IE.
     * 3. Remove the padding so developers are not caught out when they zero out
     *    "fieldset" elements in all browsers.
     */

    legend {
      box-sizing: border-box;
      /* 1 */
      color: inherit;
      /* 2 */
      display: table;
      /* 1 */
      max-width: 100%;
      /* 1 */
      padding: 0;
      /* 3 */
      white-space: normal;
      /* 1 */
    }

    /**
     * Add the correct vertical alignment in Chrome, Firefox, and Opera.
     */

    progress {
      vertical-align: baseline;
    }

    /**
     * Remove the default vertical scrollbar in IE 10+.
     */

    textarea {
      overflow: auto;
    }

    /**
     * 1. Add the correct box sizing in IE 10.
     * 2. Remove the padding in IE 10.
     */

    [type="checkbox"],
    [type="radio"] {
      box-sizing: border-box;
      /* 1 */
      padding: 0;
      /* 2 */
    }

    /**
     * Correct the cursor style of increment and decrement buttons in Chrome.
     */

    [type="number"]::-webkit-inner-spin-button,
    [type="number"]::-webkit-outer-spin-button {
      height: auto;
    }

    /**
     * 1. Correct the odd appearance in Chrome and Safari.
     * 2. Correct the outline style in Safari.
     */

    [type="search"] {
      -webkit-appearance: textfield;
      /* 1 */
      outline-offset: -2px;
      /* 2 */
    }

    /**
     * Remove the inner padding in Chrome and Safari on macOS.
     */

    [type="search"]::-webkit-search-decoration {
      -webkit-appearance: none;
    }

    /**
     * 1. Correct the inability to style clickable types in iOS and Safari.
     * 2. Change font properties to "inherit" in Safari.
     */

    ::-webkit-file-upload-button {
      -webkit-appearance: button;
      /* 1 */
      font: inherit;
      /* 2 */
    }

    /* Interactive
       ========================================================================== */

    /*
     * Add the correct display in Edge, IE 10+, and Firefox.
     */

    details {
      display: block;
    }

    /*
     * Add the correct display in all browsers.
     */

    summary {
      display: list-item;
    }

    /* Misc
       ========================================================================== */

    /**
     * Add the correct display in IE 10+.
     */

    template {
      display: none;
    }

    /**
     * Add the correct display in IE 10.
     */

    [hidden] {
      display: none;
    }
  </style>
  <style type="text/css">
    /*!
 * Milligram v1.3.0
 * https://milligram.github.io
 *
 * Copyright (c) 2017 CJ Patoilo
 * Licensed under the MIT license
 */

    *,
    *:after,
    *:before {
      box-sizing: inherit
    }

    html {
      box-sizing: border-box;
      font-size: 62.5%
    }

    body {
      color: #000000;
      font-family: 'Roboto', 'Helvetica Neue', 'Helvetica', 'Arial', sans-serif;
      font-size: 1.6em;
      font-weight: 300;
      letter-spacing: .01em;
      line-height: 1.6
    }

    blockquote {
      border-left: 0.3rem solid #d1d1d1;
      margin-left: 0;
      margin-right: 0;
      padding: 1rem 1.5rem
    }

    blockquote *:last-child {
      margin-bottom: 0
    }

    .button,
    button,
    input[type='button'],
    input[type='reset'],
    input[type='submit'] {
      background-color: #000000;
      border: 0.1rem solid #000000;
      border-radius: .4rem;
      color: #fff;
      cursor: pointer;
      display: inline-block;
      font-size: 1.1rem;
      font-weight: 700;
      height: 3.8rem;
      letter-spacing: .1rem;
      line-height: 3.8rem;
      padding: 0 3.0rem;
      text-align: center;
      text-decoration: none;
      text-transform: uppercase;
      white-space: nowrap
    }

    .button:focus,
    .button:hover,
    button:focus,
    button:hover,
    input[type='button']:focus,
    input[type='button']:hover,
    input[type='reset']:focus,
    input[type='reset']:hover,
    input[type='submit']:focus,
    input[type='submit']:hover {
      background-color: #000000;
      border-color: #000000;
      color: #fff;
      outline: 0
    }

    .button[disabled],
    button[disabled],
    input[type='button'][disabled],
    input[type='reset'][disabled],
    input[type='submit'][disabled] {
      cursor: default;
      opacity: .5
    }

    .button[disabled]:focus,
    .button[disabled]:hover,
    button[disabled]:focus,
    button[disabled]:hover,
    input[type='button'][disabled]:focus,
    input[type='button'][disabled]:hover,
    input[type='reset'][disabled]:focus,
    input[type='reset'][disabled]:hover,
    input[type='submit'][disabled]:focus,
    input[type='submit'][disabled]:hover {
      background-color: #000000;
      border-color: #000000
    }

    .button.button-outline,
    button.button-outline,
    input[type='button'].button-outline,
    input[type='reset'].button-outline,
    input[type='submit'].button-outline {
      background-color: transparent;
      color: #000000
    }

    .button.button-outline:focus,
    .button.button-outline:hover,
    button.button-outline:focus,
    button.button-outline:hover,
    input[type='button'].button-outline:focus,
    input[type='button'].button-outline:hover,
    input[type='reset'].button-outline:focus,
    input[type='reset'].button-outline:hover,
    input[type='submit'].button-outline:focus,
    input[type='submit'].button-outline:hover {
      background-color: transparent;
      border-color: #000000;
      color: #000000
    }

    .button.button-outline[disabled]:focus,
    .button.button-outline[disabled]:hover,
    button.button-outline[disabled]:focus,
    button.button-outline[disabled]:hover,
    input[type='button'].button-outline[disabled]:focus,
    input[type='button'].button-outline[disabled]:hover,
    input[type='reset'].button-outline[disabled]:focus,
    input[type='reset'].button-outline[disabled]:hover,
    input[type='submit'].button-outline[disabled]:focus,
    input[type='submit'].button-outline[disabled]:hover {
      border-color: inherit;
      color: #000000
    }

    .button.button-clear,
    button.button-clear,
    input[type='button'].button-clear,
    input[type='reset'].button-clear,
    input[type='submit'].button-clear {
      background-color: transparent;
      border-color: transparent;
      color: #000000
    }

    .button.button-clear:focus,
    .button.button-clear:hover,
    button.button-clear:focus,
    button.button-clear:hover,
    input[type='button'].button-clear:focus,
    input[type='button'].button-clear:hover,
    input[type='reset'].button-clear:focus,
    input[type='reset'].button-clear:hover,
    input[type='submit'].button-clear:focus,
    input[type='submit'].button-clear:hover {
      background-color: transparent;
      border-color: transparent;
      color: #000000
    }

    .button.button-clear[disabled]:focus,
    .button.button-clear[disabled]:hover,
    button.button-clear[disabled]:focus,
    button.button-clear[disabled]:hover,
    input[type='button'].button-clear[disabled]:focus,
    input[type='button'].button-clear[disabled]:hover,
    input[type='reset'].button-clear[disabled]:focus,
    input[type='reset'].button-clear[disabled]:hover,
    input[type='submit'].button-clear[disabled]:focus,
    input[type='submit'].button-clear[disabled]:hover {
      color: #000000
    }

    code {
      background: #f4f5f6;
      border-radius: .4rem;
      font-size: 86%;
      margin: 0 .2rem;
      padding: .2rem .5rem;
      white-space: nowrap
    }

    pre {
      background: #f4f5f6;
      border-left: 0.3rem solid #000000;
      overflow-y: hidden
    }

    pre>code {
      border-radius: 0;
      display: block;
      padding: 1rem 1.5rem;
      white-space: pre
    }

    hr {
      border: 0;
      border-top: 0.1rem solid #f4f5f6;
      margin: 3.0rem 0
    }

    input[type='email'],
    input[type='number'],
    input[type='password'],
    input[type='search'],
    input[type='tel'],
    input[type='text'],
    input[type='url'],
    textarea,
    select {
      -webkit-appearance: none;
      -moz-appearance: none;
      appearance: none;
      background-color: transparent;
      border: 0.1rem solid #d1d1d1;
      border-radius: .4rem;
      box-shadow: none;
      box-sizing: inherit;
      height: 3.8rem;
      padding: .6rem 1.0rem;
      width: 100%
    }

    input[type='email']:focus,
    input[type='number']:focus,
    input[type='password']:focus,
    input[type='search']:focus,
    input[type='tel']:focus,
    input[type='text']:focus,
    input[type='url']:focus,
    textarea:focus,
    select:focus {
      border-color: #000000;
      outline: 0
    }

    select {
      background: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" height="14" viewBox="0 0 29 14" width="29"><path fill="#d1d1d1" d="M9.37727 3.625l5.08154 6.93523L19.54036 3.625"/></svg>') center right no-repeat;
      padding-right: 3.0rem
    }

    select:focus {
      background-image: url('data:image/svg+xml;utf8,<svg xmlns="http://www.w3.org/2000/svg" height="14" viewBox="0 0 29 14" width="29"><path fill="#000000" d="M9.37727 3.625l5.08154 6.93523L19.54036 3.625"/></svg>')
    }

    textarea {
      min-height: 6.5rem
    }

    label,
    legend {
      display: block;
      font-size: 1.6rem;
      font-weight: 700;
      margin-bottom: .5rem
    }

    fieldset {
      border-width: 0;
      padding: 0
    }

    input[type='checkbox'],
    input[type='radio'] {
      display: inline
    }

    .label-inline {
      display: inline-block;
      font-weight: normal;
      margin-left: .5rem
    }

    .container {
      margin: 0 auto;
      max-width: 112.0rem;
      padding: 0 2.0rem;
      position: relative;
      width: 100%
    }

    .row {
      display: flex;
      flex-direction: column;
      padding: 0;
      width: 100%
    }

    .row.row-no-padding {
      padding: 0
    }

    .row.row-no-padding>.column {
      padding: 0
    }

    .row.row-wrap {
      flex-wrap: wrap
    }

    .row.row-top {
      align-items: flex-start
    }

    .row.row-bottom {
      align-items: flex-end
    }

    .row.row-center {
      align-items: center
    }

    .row.row-stretch {
      align-items: stretch
    }

    .row.row-baseline {
      align-items: baseline
    }

    .row .column {
      display: block;
      flex: 1 1 auto;
      margin-left: 0;
      max-width: 100%;
      width: 100%
    }

    .row .column.column-offset-10 {
      margin-left: 10%
    }

    .row .column.column-offset-20 {
      margin-left: 20%
    }

    .row .column.column-offset-25 {
      margin-left: 25%
    }

    .row .column.column-offset-33,
    .row .column.column-offset-34 {
      margin-left: 33.3333%
    }

    .row .column.column-offset-50 {
      margin-left: 50%
    }

    .row .column.column-offset-66,
    .row .column.column-offset-67 {
      margin-left: 66.6666%
    }

    .row .column.column-offset-75 {
      margin-left: 75%
    }

    .row .column.column-offset-80 {
      margin-left: 80%
    }

    .row .column.column-offset-90 {
      margin-left: 90%
    }

    .row .column.column-10 {
      flex: 0 0 10%;
      max-width: 10%
    }

    .row .column.column-20 {
      flex: 0 0 20%;
      max-width: 20%
    }

    .row .column.column-25 {
      flex: 0 0 25%;
      max-width: 25%
    }

    .row .column.column-33,
    .row .column.column-34 {
      flex: 0 0 33.3333%;
      max-width: 33.3333%
    }

    .row .column.column-40 {
      flex: 0 0 40%;
      max-width: 40%
    }

    .row .column.column-50 {
      flex: 0 0 50%;
      max-width: 50%
    }

    .row .column.column-60 {
      flex: 0 0 60%;
      max-width: 60%
    }

    .row .column.column-66,
    .row .column.column-67 {
      flex: 0 0 66.6666%;
      max-width: 66.6666%
    }

    .row .column.column-75 {
      flex: 0 0 75%;
      max-width: 75%
    }

    .row .column.column-80 {
      flex: 0 0 80%;
      max-width: 80%
    }

    .row .column.column-90 {
      flex: 0 0 90%;
      max-width: 90%
    }

    .row .column .column-top {
      align-self: flex-start
    }

    .row .column .column-bottom {
      align-self: flex-end
    }

    .row .column .column-center {
      -ms-grid-row-align: center;
      align-self: center
    }

    @media (min-width: 40rem) {
      .row {
        flex-direction: row;
        margin-left: -1.0rem;
        width: calc(100% + 2.0rem)
      }
      .row .column {
        margin-bottom: inherit;
        padding: 0 1.0rem
      }
    }

    a {
      color: #000000;
      text-decoration: none
    }

    a:focus,
    a:hover {
      color: #000000
    }

    dl,
    ol,
    ul {
      list-style: none;
      margin-top: 0;
      padding-left: 0
    }

    dl dl,
    dl ol,
    dl ul,
    ol dl,
    ol ol,
    ol ul,
    ul dl,
    ul ol,
    ul ul {
      font-size: 90%;
      margin: 1.5rem 0 1.5rem 3.0rem
    }

    ol {
      list-style: decimal inside
    }

    ul {
      list-style: circle inside
    }

    .button,
    button,
    dd,
    dt,
    li {
      margin-bottom: 1.0rem
    }

    fieldset,
    input,
    select,
    textarea {
      margin-bottom: 1.5rem
    }

    blockquote,
    dl,
    figure,
    form,
    ol,
    p,
    pre,
    table,
    ul {
      margin-bottom: 2.5rem
    }

    table {
      border-spacing: 0;
      width: 100%
    }

    td,
    th {
      border-bottom: 0.1rem solid #e1e1e1;
      padding: 1.2rem 1.5rem;
      text-align: left
    }

    td:first-child,
    th:first-child {
      padding-left: 0
    }

    td:last-child,
    th:last-child {
      padding-right: 0
    }

    b,
    strong {
      font-weight: bold
    }

    p {
      margin-top: 0
    }

    h1,
    h2,
    h3,
    h4,
    h5,
    h6 {
      font-weight: 300;
      letter-spacing: -.1rem;
      margin-bottom: 2.0rem;
      margin-top: 0
    }

    h1 {
      font-size: 4.6rem;
      line-height: 1.2
    }

    h2 {
      font-size: 3.6rem;
      line-height: 1.25
    }

    h3 {
      font-size: 2.8rem;
      line-height: 1.3
    }

    h4 {
      font-size: 2.2rem;
      letter-spacing: -.08rem;
      line-height: 1.35
    }

    h5 {
      font-size: 1.8rem;
      letter-spacing: -.05rem;
      line-height: 1.5
    }

    h6 {
      font-size: 1.6rem;
      letter-spacing: 0;
      line-height: 1.4
    }

    img {
      max-width: 100%
    }

    .clearfix:after {
      clear: both;
      content: ' ';
      display: table
    }

    .float-left {
      float: left
    }

    .float-right {
      float: right
    }
  </style>
</body>

</html>`
