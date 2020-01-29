import os
from flask import Flask
from flask_restful import reqparse, abort, Api, Resource

import ledGrid

app = Flask(__name__)
api = Api(app)

IMAGES = [x.split('.')[0] for x in filter(lambda x: x.endswith('.png'), os.listdir('images'))]
EFFECTS = ['scroll_left', 'scroll_right', 'type']

parser = reqparse.RequestParser()
# parser.add_argument('name')


class Images(Resource):
    def get(self):
        return IMAGES

class Image(Resource):
    def post(self, name):
        if name in IMAGES:
            ledGrid.image_display(name)
            return name, 201
        if len(name) == 64 and int(name) >= 0:
            ledGrid.string_display(name)
            return name, 201
        else:
            abort(404, message=f"{name} does not exist")

class Text(Resource):
    def post(self, effect):
        if effect in EFFECTS:
            return effect, 201
        else:
            abort(404, message=f"{effect} does not exist")


api.add_resource(Images, '/images')
api.add_resource(Image, '/image/<name>')
api.add_resource(Text, '/text/<effect>')


if __name__ == '__main__':
    app.run(debug=True)
