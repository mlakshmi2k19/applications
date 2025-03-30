import 'package:flutter/material.dart';
import 'package:video_shorts/videoplayer_remote.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Shorts',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        primarySwatch: Colors.teal,
      ),
      home: const MyHomePage(title: 'Shorts'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key, required this.title}) : super(key: key);

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  PageController controller = PageController();
  static dynamic currentPageValue = 0.0;

  List pageViewItem = [
    page('https://filesamples.com/samples/video/mp4/sample_640x360.mp4'),
    page('http://techslides.com/demos/sample-videos/small.mp4'),
    page(
        'https://archive.org/download/SampleVideo1280x7205mb/SampleVideo_1280x720_5mb.mp4'),
    page(
        'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4'),
    page(
        'https://archive.org/download/SampleVideo1280x7205mb/SampleVideo_1280x720_5mb.mp4'),
    page(
        'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4'),
    page('https://filesamples.com/samples/video/mp4/sample_640x360.mp4'),
    page('http://techslides.com/demos/sample-videos/small.mp4'),
    page(
        'https://archive.org/download/SampleVideo1280x7205mb/SampleVideo_1280x720_5mb.mp4'),
    page(
        'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4'),
    page('https://filesamples.com/samples/video/mp4/sample_640x360.mp4'),
    page('http://techslides.com/demos/sample-videos/small.mp4'),
    page(
        'https://archive.org/download/SampleVideo1280x7205mb/SampleVideo_1280x720_5mb.mp4'),
    page('https://filesamples.com/samples/video/mp4/sample_640x360.mp4'),
    page('http://techslides.com/demos/sample-videos/small.mp4'),
    page(
        'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4'),
    page(
        'https://archive.org/download/SampleVideo1280x7205mb/SampleVideo_1280x720_5mb.mp4'),
    page('https://filesamples.com/samples/video/mp4/sample_640x360.mp4'),
    page('http://techslides.com/demos/sample-videos/small.mp4'),
    page(
        'http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ElephantsDream.mp4'),
    
  ];

  @override
  void initState() {
    super.initState();
    controller.addListener(() {
      setState(() {
        currentPageValue = controller.page;
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // appBar: AppBar(
      //   title: Text(widget.title),
      // ),
      body: PageView.builder(
          itemCount: pageViewItem.length,
          scrollDirection: Axis.vertical,
          controller: controller,
          itemBuilder: (context, position) {
            return Transform(
              transform: Matrix4.identity()
                ..rotateX(currentPageValue - position),
              child: pageViewItem[position],
            );
          }),
    );
  }
}

Widget page(String url) {
  // ignore: avoid_unnecessary_containers, sized_box_for_whitespace
  return Container(
    width: double.infinity,
    height: double.infinity,
    child: VideoPlayerRemote(
      url: url,
    ),
  );
}
